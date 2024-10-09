package bc833

import (
	"github.com/stretchr/testify/assert"
	"temnok/pcbc/eda"
	"testing"
)

func Test_BC833(t *testing.T) {
	pcb := eda.NewPCB(36, 46)

	pcb.Component(Board)

	assert.NoError(t, pcb.SaveFiles("out/bc833-"))
}
