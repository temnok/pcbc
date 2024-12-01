package oc

import "io"

// Write passes an OC-token (string data ended with a curly brace) to
// the provided writer, escaping curly braces within the data with '\'.
// If '\' is the last data symbol then ' ' (space) is added after it.
// Any error from the writer is returned immediately.
func Write(w io.Writer, data []byte, oc bool) error {
	i0 := 0

	// write data with escaped curly braces
	for i, b := range data {
		if b != symOpen && b != symClose {
			continue
		}

		if i0 < i {
			if _, err := w.Write(data[i0:i]); err != nil {
				return err
			}
		}

		if _, err := w.Write(sliceEscape); err != nil {
			return err
		}

		i0 = i
	}

	// write remaining data
	if n := len(data); i0 < n {
		if _, err := w.Write(data[i0:]); err != nil {
			return err
		}

		// if escape is the last data symbol then add a space
		if data[n-1] == symEscape {
			if _, err := w.Write(sliceSpace); err != nil {
				return err
			}
		}
	}

	// write a curly brace
	if oc {
		_, err := w.Write(sliceOpen)
		return err
	} else {
		_, err := w.Write(sliceClose)
		return err
	}
}
