// Copyright © 2025 Alex Temnok. All rights reserved.

package ws2812b_2020

import (
	"temnok/pcbc/eda/pcb"
	"temnok/pcbc/eda/pcb/config"
	"testing"
)

func TestBoard(t *testing.T) {
	conf := config.Default()
	conf.Width, conf.Height = 12.5, 5

	pcb.Process(conf, Board)
}
