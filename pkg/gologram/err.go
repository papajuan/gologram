package gologram

/**
 * @author  papajuan
 * @date    1/4/2025
 **/

type Err struct {
	err error
}

func NewErr(err error) *Err {
	return &Err{err: err}
}
