// Copyright © 2025 Alex Temnok. All rights reserved.

package lbrn

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_f2s(t *testing.T) {
	assert.Equal(t, "1", f2s(1))
	assert.Equal(t, "-1.23", f2s(-1.23))

	assert.Equal(t, "0", f2s(0.0000000001))
	assert.Equal(t, "0", f2s(0.0000000004))
	assert.Equal(t, "0.000000001", f2s(0.0000000005))
	assert.Equal(t, "0.000000001", f2s(0.0000000009))

	assert.Equal(t, "-0", f2s(-0.0000000001))
	assert.Equal(t, "-0", f2s(-0.0000000004))
	assert.Equal(t, "-0.000000001", f2s(-0.0000000005))
	assert.Equal(t, "-0.000000001", f2s(-0.0000000009))

	assert.Equal(t, "0.999999999", f2s(0.999999999))
	assert.Equal(t, "1000000000", f2s(999999999.999999999))
	assert.Equal(t, "1000000000000000000", f2s(999_999_999_999_999_999))
}
