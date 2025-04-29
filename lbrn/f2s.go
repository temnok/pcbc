package lbrn

import "strconv"

// f2s converts float number to a string performing rounding to 9 decimal places.
// Having this method is important because default Sprint(f)
// will result in different outputs on AMD and ARM platforms
func f2s(val float64) string {
	str := strconv.FormatFloat(val, 'f', 9, 64)

	for str[len(str)-1] == '0' {
		str = str[:len(str)-1]
	}

	if str[len(str)-1] == '.' {
		str = str[:len(str)-1]
	}

	return str
}
