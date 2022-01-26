package functions

import "strconv"

// Format a number with grouped thousands
//
// Actually the sypnosis of NumberFormat() is like this:
//
//  NumberFormat(s float64[, decimals int[, decimal_separator string[, thousands_separator string]]])
func NumberFormat(num float64, args ...interface{}) string {
	var decimals int
	var decimal_separator string
	var thousands_separator string

	// arguments
	if len(args) > 0 {
		decimals = args[0].(int)
	} else {
		decimals = 0
	}
	if len(args) > 1 {
		decimal_separator = args[1].(string)
	} else {
		decimal_separator = "."
	}
	if len(args) > 2 {
		thousands_separator = args[2].(string)
	} else {
		thousands_separator = ","
	}
	// negative
	if num < 0 {
		return "-" + NumberFormat(-num, decimals, decimal_separator, thousands_separator)
	}

	// round and convert float to string
	s := strconv.FormatFloat(num, 'f', decimals, 64)

	// culculate length and sep num
	s_len := len(s)
	i_len := 0   // length of integer part
	sep_num := 0 // number of separators

	tsep_len := len(thousands_separator)
	dsep_len := len(decimal_separator)

	for i := 0; i < s_len; i++ {
		if s[i] == '.' {
			break
		}
		i_len++
	}
	sep_num = int(i_len / 3)

	// copy
	dst := make([]byte, s_len+sep_num*tsep_len+dsep_len)
	pos := 0
	for i := 0; i < s_len; i++ {

		if i < i_len {
			// integer part
			if i_len%3 == i%3 {
				for j := 0; j < tsep_len; j++ {
					dst[pos] = thousands_separator[j]
					pos++
				}
			}
		} else if i == i_len {
			// decimal part
			for j := 0; j < dsep_len; j++ {
				dst[pos] = decimal_separator[j]
				pos++
			}
			continue
		}
		dst[pos] = s[i]
		pos++
	}
	return string(dst[0:pos])
}
