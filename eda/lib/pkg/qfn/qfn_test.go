package qfn

import (
	"github.com/stretchr/testify/assert"
	"temnok/pcbc/eda"
	"testing"
)

func TestQFN16(t *testing.T) {
	assert.NoError(t, eda.NewPCB(QFN16G).SaveOverview())
}
