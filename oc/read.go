// Copyright © 2025 Alex Temnok. All rights reserved.

package oc

import "io"

// Read obtains an OC-token (string data ended with a curly brace) from
// the provided reader, un-escaping curly braces within the data if necessary.
// Any error from the reader is returned immediately.
func Read(r io.Reader) ([]byte, bool, error) {
	data := []byte{0}

	for i := 0; ; {
		if n, err := r.Read(data[i:]); err != nil {
			return nil, false, err
		} else if n == 0 { // handle unlikely case mentioned in io.Reader.Read's contract
			continue
		}

		if data[i] == symOpen || data[i] == symClose {
			// Un-escape if curly brace was escaped
			if i > 0 && data[i-1] == symEscape {
				data[i-1] = data[i]
				continue
			}
			
			return data[i:], data[i] == symOpen, nil
		}

		data = append(data, 0)
		i++
	}
}
