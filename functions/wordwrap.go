package functions

import (
	"errors"
)

func isSpace(c byte) bool {
	//chars = " \n\r\t\v\x00"
	if c == ' ' || c == '\n' || c == '\r' || c == '\t' || c == 0x00 || c == '\v' {
		return true
	}
	return false
}

// Wraps a string to a given number of characters
//
// Actually the sypnosis of Wordwrap() is like this:
//
//  Wordwrap(s string[, width int = 75[, break string = "\n", [ cut_long_words bool = false ]]])
func Wordwrap(s string, args ...interface{}) (string, error) {
	var width int = 75
	var break_s string = "\n"
	var cut_long_words bool = false

	// arguments
	args_len := len(args)
	if args_len >= 1 {
		switch tmp := args[0].(type) {
		case int:
			width = tmp
		default:
			return "", errors.New("width must be int")
		}
	}
	if args_len >= 2 {
		switch tmp := args[1].(type) {
		case string:
			break_s = tmp
		default:
			return "", errors.New("break_s must be string")
		}
	}
	if args_len >= 3 {
		switch tmp := args[2].(type) {
		case bool:
			cut_long_words = tmp
		default:
			return "", errors.New("cut_long_words must be string")
		}
	}
	_ = cut_long_words

	// walk chars
	s_len := len(s)
	pos := 0
	line_pos := 0

	dst := []byte("")
	var space_c int16 = -1

	for pos < s_len {
		pos_e := pos

		if !isSpace(s[pos_e]) {
			pos_e++

			for pos_e < s_len && !isSpace(s[pos_e]) {
				pos_e++
			}
		} else {
			pos_e++
			for pos_e < s_len && isSpace(s[pos_e]) {
				pos_e++
			}
		}
		if (pos_e - line_pos) > width {
			if cut_long_words {
				for pos_e-pos > width {
					dst = append(dst, break_s...)
					dst = append(dst, s[pos:pos+width]...)
					pos += width
				}
			}

			dst = append(dst, break_s...)
			dst = append(dst, s[pos:pos_e]...)
			if pos_e < s_len {
				space_c = int16(s[pos_e])
			}
			line_pos = pos + 1
			pos = pos_e + 1
		} else {
			if space_c >= 0 {
				dst = append(dst, byte(space_c))
			}
			dst = append(dst, s[pos:pos_e]...)
			if pos_e < s_len {
				space_c = int16(s[pos_e])
			}
			pos = pos_e + 1
		}
	}
	return string(dst), nil
}
