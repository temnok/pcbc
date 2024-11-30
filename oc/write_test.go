package oc

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type dummyWriter struct {
	strings []string
}

func (w *dummyWriter) Write(p []byte) (int, error) {
	w.strings = append(w.strings, string(p))
	return len(p), nil
}

func TestWrite(t *testing.T) {
	tests := []struct {
		data     []byte
		oc       bool
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
	}

	for _, test := range tests {
		w := &dummyWriter{}
		assert.NoError(t, Write(w, test.data, test.oc))
		assert.Equal(t, test.expected, w.strings)
	}
}
