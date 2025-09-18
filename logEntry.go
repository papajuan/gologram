package gologram

/**
 * @author  papajuan
 * @date    9/18/2025
 **/

type logEntry struct {
	data  []byte
	flush chan struct{}
}
