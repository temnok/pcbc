package util

func GoAll(funcs []func() error) error {
	res := make(chan error)

	for _, fn := range funcs {
		go func() {
			res <- fn()
		}()
	}

	var firstErr error

	for cnt := 0; cnt < len(funcs); cnt++ {
		if err := <-res; err != nil && firstErr == nil {
			firstErr = err
		}
	}

	return firstErr
}
