package util

import "errors"

// RunConcurrently runs provided functions concurrently, blocking until all functions are completed.
// It returns resulting errors joined into the single one or nil if functions returned no errors.
func RunConcurrently(funs []func() error) error {
	wait := make(chan error)

	for _, fun := range funs {
		go func() {
			wait <- fun()
		}()
	}

	var errs []error

	for cnt := 0; cnt < len(funs); cnt++ {
		errs = append(errs, <-wait)
	}

	return errors.Join(errs...)
}
