package gologram

import (
	"gologram/buffer"
	"sync"
)

/**
 * @author  papajuan
 * @date    1/4/2025
 **/

func Sync() error {
	var wg sync.WaitGroup
	wg.Add(2)
	var res error
	go func() {
		err := buffer.Stdout().Sync()
		if err != nil {
			res = err
		}
	}()
	go func() {
		err := buffer.Stderr().Sync()
		if err != nil {
			res = err
		}
	}()
	wg.Wait()
	return res
}
