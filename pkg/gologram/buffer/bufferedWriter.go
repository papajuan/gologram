package buffer

import (
	"bufio"
	"errors"
	"os"
	"sync"
	"time"
)

/**
 * @author  papajuan
 * @date    1/4/2025
 **/

var (
	stdoutInstance *BufferedWriter
	stderrInstance *BufferedWriter
	onceOut        sync.Once
	onceErr        sync.Once
	defaultSize    = 512             // Default buffer size (512 B)
	defaultFlush   = 5 * time.Second // Default flush interval
)

// BufferedWriter is a thread-safe buffered writer for os output.
type BufferedWriter struct {
	mu            sync.Mutex
	writer        *bufio.Writer
	bufferSize    int
	flushInterval time.Duration
	ticker        *time.Ticker
	stopChan      chan bool // Signals the background flusher to stop
	doneChan      chan bool // Signals that the flusher has stopped
}

func Initialize() {
	go Stdout()
	go Stderr()
}

// Stdout returns the singleton instance of BufferedWriter for Stdout.
func Stdout() *BufferedWriter {
	onceOut.Do(func() {
		stdoutInstance = newBufferedWriter(defaultSize, defaultFlush, os.Stdout)
	})
	return stdoutInstance
}

// Stderr returns the singleton instance of BufferedWriter for Stderr.
func Stderr() *BufferedWriter {
	onceErr.Do(func() {
		stderrInstance = newBufferedWriter(defaultSize, defaultFlush, os.Stderr)
	})
	return stderrInstance
}

// newBufferedWriter creates a new BufferedWriter with defaults if unspecified.
func newBufferedWriter(bufferSize int, flushInterval time.Duration, output *os.File) *BufferedWriter {
	if bufferSize <= 0 {
		bufferSize = defaultSize
	}
	if flushInterval <= 0 {
		flushInterval = defaultFlush
	}
	bsw := &BufferedWriter{
		writer:        bufio.NewWriterSize(output, bufferSize),
		bufferSize:    bufferSize,
		flushInterval: flushInterval,
		stopChan:      make(chan bool, 1), // Buffered to prevent blocking
		doneChan:      make(chan bool),    // Unbuffered for signal safety
	}
	go bsw.backgroundFlusher()
	return bsw
}

// Write writes data to the buffer, flushing if the buffer is full.
func (bsw *BufferedWriter) Write(p []byte) (int, error) {
	bsw.mu.Lock()
	defer bsw.mu.Unlock()
	n, err := bsw.writer.Write(p)
	if err != nil {
		return n, err
	}
	// Flush if the buffer is full
	if bsw.writer.Buffered() >= bsw.bufferSize {
		err = bsw.writer.Flush()
	}
	return n, err
}

// flush writes the contents of the buffer to stdout.
func (bsw *BufferedWriter) flush() error {
	bsw.mu.Lock()
	defer bsw.mu.Unlock()
	return bsw.writer.Flush()
}

// backgroundFlusher periodically flushes the buffer.
func (bsw *BufferedWriter) backgroundFlusher() {
	defer close(bsw.doneChan) // Signal that the flusher has stopped
	bsw.ticker = time.NewTicker(bsw.flushInterval)
	defer bsw.ticker.Stop()
	for {
		select {
		case <-bsw.ticker.C:
			_ = bsw.flush() // Flush periodically
		case <-bsw.stopChan:
			_ = bsw.flush() // Final flush before stopping
			return
		}
	}
}

// Sync flushes the buffer and stops the background flusher.
func (bsw *BufferedWriter) Sync() error {
	if bsw.stopChan != nil {
		select {
		case bsw.stopChan <- true: // Signal the flusher to stop
		case <-time.After(5 * time.Second): // Prevent blocking forever
			return errors.New("failed to stop the background flusher: timeout")
		}
	}
	select {
	case <-bsw.doneChan: // Wait for the flusher to stop
	case <-time.After(5 * time.Second): // Prevent blocking forever
		return errors.New("timeout waiting for the flusher to stop")
	}
	return bsw.flush() // Perform a final flush
}
