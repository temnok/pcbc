// Copyright © 2025 Alex Temnok. All rights reserved.

package util

func Repeat[T any](val T, n int) []T {
	arr := make([]T, n)

	for i := range n {
		arr[i] = val
	}

	return arr
}
