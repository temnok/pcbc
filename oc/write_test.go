package oc

import (
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type dummyWriter struct {
	strings            []string
	stringsLen, errPos int
}

func (w *dummyWriter) Write(p []byte) (int, error) {
	w.strings = append(w.strings, string(p))

	if w.stringsLen += len(p); w.errPos > 0 && w.errPos <= w.stringsLen {
		return 0, errors.New("test")
	}

	return len(p), nil
}

func TestWrite(t *testing.T) {
	tests := []struct {
		data     []byte
		oc       bool
		errPos   int
		expected []string
	}{
		{
			data:     nil,
			oc:       true,
			expected: []string{"{"},
		},

		{
			data:     nil,
			oc:       false,
			expected: []string{"}"},
		},

		{
			data:     []byte("key"),
			oc:       true,
			expected: []string{"key", "{"},
		},

		{
			data:     []byte("value"),
			oc:       false,
			expected: []string{"value", "}"},
		},

		{
			data:     []byte("key with { and \\} ended with \\"),
			oc:       true,
			expected: []string{"key with ", "\\", "{ and \\", "\\", "} ended with \\", " ", "{"},
		},

		{
			data:     []byte("value with { and \\} ended with \\"),
			oc:       false,
			expected: []string{"value with ", "\\", "{ and \\", "\\", "} ended with \\", " ", "}"},
		},

		{
			data:   []byte("e{"),
			errPos: 1,
		},

		{
			data:   []byte("e{"),
			errPos: 2,
		},

		{
			data:   []byte("e"),
			errPos: 1,
		},

		{
			data:   []byte("\\"),
			errPos: 2,
		},
	}

	for _, test := range tests {
		w := &dummyWriter{errPos: test.errPos}

		if test.errPos > 0 {
			assert.Error(t, Write(w, test.data, test.oc))
			continue
		}

		assert.NoError(t, Write(w, test.data, test.oc))
		assert.Equal(t, test.expected, w.strings,
			fmt.Sprintf("data=%q, oc=%v, errPos=%v\n", test.data, test.oc, test.errPos))
	}
}
