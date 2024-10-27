package x2tiny

import (
	"github.com/stretchr/testify/assert"
	"temnok/pcbc/eda"
	"testing"
)

func TestDemo(t *testing.T) {
	assert.NoError(t, eda.GeneratePCB(X2Base))
}
