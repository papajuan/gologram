package gologram

/**
 * @author  papajuan
 * @date    1/4/2025
 **/

type Format uint8

const (
	CONSOLE Format = iota
	JSON
)

func NewLogFormat(f string) Format {
	switch f {
	case "JSON":
		return JSON
	default:
		return CONSOLE
	}
}
