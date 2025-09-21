package ptr

func To[T any](val T) *T {
	return &val
}

func Val[T any](ptr *T) T {
	if ptr == nil {
		var zero T
		return zero
	}

	return *ptr
}
