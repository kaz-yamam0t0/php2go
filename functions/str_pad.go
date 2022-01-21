/**
 * Golang equivalent to php `str_pad`
 *
 * Pad a string to a certain length with another string
 * @see https://www.php.net/manual/en/function.str-pad.php
 *
 * @param string s
 * @param int length
 * @param string pad_string
 * @param int pad_type
 * @return string
 */
package functions

func StrPad(s string, length int, args ...interface{}) string {
	var pad_string string
	pad_type := STR_PAD_RIGHT

	// default arguments
	if len(args) > 0 {
		pad_string = args[0].(string)
	} else {
		pad_string = " "
	}
	if len(args) > 1 {
		pad_type = args[1].(int)
	}

	// culculate
	s_len := len(s)
	if s_len >= length {
		return s[:]
	}
	n := length - s_len
	left := 0
	right := 0

	switch pad_type {
	case STR_PAD_LEFT:
		left = n
	case STR_PAD_BOTH:
		left = int(n / 2)
		right = n - left
	default:
		right = n
	}

	pad_str_len := len(pad_string)

	// strcpy
	dst := make([]byte, length)
	pos := 0
	for i := 0; i < left; i++ {
		dst[pos] = pad_string[i%pad_str_len]
		pos++
	}
	for i := 0; i < s_len; i++ {
		dst[pos] = s[i]
		pos++
	}
	for i := 0; i < right; i++ {
		dst[pos] = pad_string[i%pad_str_len]
		pos++
	}
	return string(dst)
}
