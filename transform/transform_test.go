// Copyright © 2025 Alex Temnok. All rights reserved.

package transform

import (
	"fmt"
	"strings"
	"testing"
)

func TestAscii(t *testing.T) {
	tests := []struct {
		input     []string
		transform T
		expected  []string
	}{
		0: {
			input: []string{
				"12",
				"3.",
			},
			transform: Move(3, 2),
			expected: []string{
				".....",
				".....",
				"...12",
				"...3.",
			},
		},
		1: {
			input: []string{
				"12",
				"3.",
			},
			transform: Scale(3, 2),
			expected: []string{
				"111222",
				"111222",
				"333...",
				"333...",
			},
		},
		2: {
			input: []string{
				"12",
				"3.",
			},
			transform: Rotate(90).Move(2, 0),
			expected: []string{
				"31",
				".2",
			},
		},
		3: {
			input: []string{
				"12",
				"3.",
			},
			transform: Rotate(-90).Move(0, 2).Scale(2, 3).Move(3, 2),
			expected: []string{
				".......",
				".......",
				"...22..",
				"...22..",
				"...22..",
				"...1133",
				"...1133",
				"...1133",
			},
		},
	}

	for testID, test := range tests {
		t.Run(fmt.Sprintf("Test #%v", testID), func(t *testing.T) {

			buf := [][]byte{{}}
			for y, row := range test.input {
				for x, val := range row {
					if val == ' ' {
						continue
					}
					ax, ay := test.transform.Apply(float64(x), float64(y))
					bx, by := test.transform.Apply(float64(x+1), float64(y+1))
					x0, y0 := int(min(ax, bx)), int(min(ay, by))
					x1, y1 := int(max(ax, bx)), int(max(ay, by))
					for len(buf) < y1 {
						newRow := make([]byte, len(buf[0]))
						for i := range newRow {
							newRow[i] = '.'
						}
						buf = append(buf, newRow)
					}
					for i := range buf {
						for len(buf[i]) < x1 {
							buf[i] = append(buf[i], '.')
						}
					}
					for i := y0; i < y1; i++ {
						for j := x0; j < x1; j++ {
							buf[i][j] = byte(val)
						}
					}
				}
			}
			var actual []byte
			for _, row := range buf {
				actual = append(actual, row...)
				actual = append(actual, '\n')
			}

			if want, got := strings.Join(test.expected, "\n")+"\n", string(actual); want != got {
				t.Errorf("want: %v\n got: %v\n", want, got)
			}
		})
	}
}
