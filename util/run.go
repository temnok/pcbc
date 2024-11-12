package util

import "errors"

// RunConcurrently runs provided functions concurrently, blocking until all functions are completed.
// It returns functions errors joined into the single one or nil if functions returned no errors.
func RunConcurrently(funcs []func() error) error {
	wait := make(chan error)

	for _, fn := range funcs {
		go func() {
			wait <- fn()
		}()
	}

	var errs []error

	for cnt := 0; cnt < len(funcs); cnt++ {
		errs = append(errs, <-wait)
	}

	return errors.Join(errs...)
}
