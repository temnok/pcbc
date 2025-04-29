// Copyright © 2025 Alex Temnok. All rights reserved.

package util

import "errors"

// RunConcurrently runs provided functions concurrently, blocking until all functions are completed.
// It returns resulting errors joined into the single one or nil if functions returned no errors.
func RunConcurrently(funcs ...func() error) error {
	wait := make(chan error)

	for _, f := range funcs {
		go func() {
			wait <- f()
		}()
	}

	var errs []error

	for range len(funcs) {
		errs = append(errs, <-wait)
	}

	return errors.Join(errs...)
}
