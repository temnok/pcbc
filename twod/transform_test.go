package twod

import (
	"fmt"
	"math"
	"strings"
	"testing"
)

const degree = math.Pi / 180

func TestAscii(t *testing.T) {
	tests := []struct {
		input     []string
		transform Transform
		expected  []string
	}{
		0: {
			input: []string{
				"12",
				"3.",
			},
			transform: Move(Coord{3, 2}),
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
			transform: Scale(Coord{3, 2}),
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
			transform: Move(Coord{2, 0}).Rotate(90 * degree),
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
			transform: Move(Coord{3, 2}).Scale(Coord{2, 3}).Move(Coord{0, 2}).Rotate(-90 * degree),
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
					a := test.transform.Point(Coord{float64(x), float64(y)})
					b := test.transform.Point(Coord{float64(x + 1), float64(y + 1)})
					x0, y0 := int(min(a.X, b.X)), int(min(a.Y, b.Y))
					x1, y1 := int(max(a.X, b.X)), int(max(a.Y, b.Y))
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
			actual := []byte{}
			for _, row := range buf {
				actual = append(actual, row...)
				actual = append(actual, '\n')
			}

			if got, want := string(actual), strings.Join(test.expected, "\n")+"\n"; got != want {
				t.Errorf("Got:\n%v\nWant:\n%v\n", got, want)
			}
		})
	}
}
