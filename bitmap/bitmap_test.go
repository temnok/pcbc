// Copyright © 2025 Alex Temnok. All rights reserved.

package bitmap

import (
	"testing"
)

func TestBitmap_Mask(t *testing.T) {
	tests := []struct {
		want uint64
		i, j int
	}{
		{0, 1, 1},
		{1, 0, 1},
		{3, 0, 2},
		{7, 0, 3},
		{15, 0, 4},
		{0x18, 3, 5},
		{0x7FFF_FFFF_FFFF_FFFF, 0, 63},
		{0xFFFF_FFFF_FFFF_FFFE, 1, 64},
		{0xFFFF_FFFF_FFFF_FFFF, 0, 64},
		{0xFFFF_FFFF_FFFF_FFFF, 64, 0},
	}

	for _, test := range tests {
		got := mask(test.i, test.j)
		if test.want != got {
			t.Errorf("mask(%d, %d):\nwant %x\n got %x\n", test.i, test.j, test.want, got)
		}
	}
}
