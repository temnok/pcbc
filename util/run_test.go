// Copyright © 2025 Alex Temnok. All rights reserved.

package util

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestRunConcurrently(t *testing.T) {
	errs := []error{
		errors.New("zero"),
		errors.New("one"),
		errors.New("two"),
	}

	tests := []struct {
		name     string
		input    []func() error
		expected []error
	}{
		{
			name: "no errors",
			input: []func() error{
				func() error { return nil },
				func() error { return nil },
				func() error { return nil },
			},
			expected: nil,
		},

		{
			name: "all errors",
			input: []func() error{
				func() error { time.Sleep(1 * time.Millisecond); return errs[0] },
				func() error { time.Sleep(10 * time.Millisecond); return errs[1] },
				func() error { time.Sleep(20 * time.Millisecond); return errs[2] },
			},
			expected: errs,
		},

		{
			name: "some errors",
			input: []func() error{
				func() error { time.Sleep(3 * time.Millisecond); return errs[0] },
				func() error { time.Sleep(2 * time.Millisecond); return nil },
				func() error { time.Sleep(1 * time.Millisecond); return errs[2] },
			},
			expected: []error{errs[2], errs[0]},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := RunConcurrently(test.input...)
			if err == nil {
				assert.Nil(t, test.expected)
			} else {
				assert.ElementsMatch(t, test.expected, err.(interface{ Unwrap() []error }).Unwrap())
			}
		})
	}
}
