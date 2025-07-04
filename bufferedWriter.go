package gologram

import (
	"bufio"
	"os"
	"sync"
	"sync/atomic"
	"time"
)

/**
 * @author  papajuan
 * @date    1/4/2025
 **/

var (
	stdoutInstance *BufferedWriter
	onceOut        sync.Once
	stderrInstance *BufferedWriter
	onceErr        sync.Once
	defaultFlush   = 1 * time.Second // Default flush interval

	// Global buffer pools for small and large logs
	smallBufferPool = sync.Pool{
		New: func() interface{} {
			return make([]byte, 1024) // 1KB buffer (standard log size)
		},
	}
	largeBufferPool = sync.Pool{
		New: func() interface{} {
			return make([]byte, 8192) // 8KB buffer for large logs
		},
	}
)

// BufferedWriter is a high-performance buffered writer inspired by Uber Zap.
type BufferedWriter struct {
	writer   *bufio.Writer
	logQueue chan []byte
	stopChan chan struct{}
	closed   int32 // Atomic flag to check if logQueue is closed
	once     sync.Once
}

// initialize stdout instance
func initialize() {
	stdout()
	stderr()
}

// stdout returns the singleton instance of BufferedWriter for stdout.
func stdout() *BufferedWriter {
	onceOut.Do(func() {
		stdoutInstance = newBufferedWriter(os.Stdout)
	})
	return stdoutInstance
}

// stderr returns the singleton instance of BufferedWriter for stderr.
func stderr() *BufferedWriter {
	onceErr.Do(func() {
		stderrInstance = newBufferedWriter(os.Stderr)
	})
	return stderrInstance
}

// newBufferedWriter creates a new BufferedWriter optimized for high performance.
func newBufferedWriter(output *os.File) *BufferedWriter {
	bsw := &BufferedWriter{
		writer:   bufio.NewWriterSize(output, 4096), // 4KB buffer
		logQueue: make(chan []byte, 10000),          // Large buffer for async logging
		stopChan: make(chan struct{}),
		closed:   0, // Not closed
	}
	go bsw.processLogQueue()
	go bsw.backgroundFlusher()
	return bsw
}

// Write asynchronously enqueues logs for batch processing (zero-mutex).
func (bsw *BufferedWriter) Write(p []byte) (int, error) {
	// Prevent sending to closed channel
	if atomic.LoadInt32(&bsw.closed) == 1 {
		return 0, nil // Ignore writes after shutdown
	}

	// Get buffer from pool based on log size
	var buf []byte
	if len(p) <= 1024 {
		buf = smallBufferPool.Get().([]byte)[:len(p)]
	} else if len(p) <= 8192 {
		buf = largeBufferPool.Get().([]byte)[:len(p)]
	} else {
		buf = make([]byte, len(p)) // Allocate only when necessary
	}
	copy(buf, p) // Copy log entry
	// Enqueue log message (non-blocking)
	select {
	case bsw.logQueue <- buf:
	default:
		// Drop log if queue is full (prevents deadlock)
	}

	return len(buf), nil
}

// processLogQueue continuously writes logs from the channel to bufio.Writer.
func (bsw *BufferedWriter) processLogQueue() {
	for log := range bsw.logQueue {
		if bsw.writer.Available() < len(log) {
			bsw.writer.Flush() // Flush only when buffer is full
		}
		bsw.writer.Write(log)
		// Return buffer to pool properly
		if cap(log) == 1024 {
			smallBufferPool.Put(log[:1024]) // Reset before returning
		}
	}
}

// Sync flushes the buffer and stops background logging.
func (bsw *BufferedWriter) Sync() error {
	bsw.once.Do(func() {
		// Mark as closed
		atomic.StoreInt32(&bsw.closed, 1)

		close(bsw.stopChan) // Stop background flusher

		// Drain remaining logs before closing the queue
		go func() {
			for log := range bsw.logQueue {
				bsw.writer.Write(log)
				switch cap(log) {
				case 1024:
					smallBufferPool.Put(log[:1024]) // Reset buffer before returning
				case 8192:
					largeBufferPool.Put(log[:8192]) // Reset buffer before returning
				}
			}
			bsw.writer.Flush()
		}()

		close(bsw.logQueue) // Now safe to close the channel
	})

	return bsw.flush()
}

// flush ensures all buffered logs are written to stdout.
func (bsw *BufferedWriter) flush() error {
	return bsw.writer.Flush()
}

// backgroundFlusher periodically flushes the buffer.
func (bsw *BufferedWriter) backgroundFlusher() {
	ticker := time.NewTicker(defaultFlush)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			_ = bsw.flush()
		case <-bsw.stopChan:
			_ = bsw.flush() // Final flush before stopping
			return
		}
	}
}
