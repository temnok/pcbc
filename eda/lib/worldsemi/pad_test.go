package worldsemi

import (
	"github.com/stretchr/testify/assert"
	"temnok/pcbc/eda"
	"testing"
)

func TestBoard(t *testing.T) {
	pcb := eda.NewPCB(5, 5, WS2812B_2020)

	assert.NoError(t, pcb.SaveFiles("gen/"))
}
