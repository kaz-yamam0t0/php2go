/**
 * Golang equivalent to php `ltrim`
 *
 * Strip whitespace (or other characters) from the beginning of a string
 * @see https://www.php.net/manual/en/function.ltrim.php
 *
 * @param string s
 * @param string characters
 * @return string
 */
package functions

func Ltrim(s string, args ...string) string {
	var flags [256]byte

	var chars string
	if len(args) > 0 {
		chars = args[0]
	} else {
		chars = " \n\r\t\v\x00"
	}

	// flags
	len_ := len(chars)
	for i := 0; i < len_; i++ {
		c := chars[i]
		if len_ > i+2 && chars[i+1] == '.' && chars[i+2] == '.' {
			if len_ <= i+3 {
				// Invalid '..'-range, '..'-range needs to be incrementing
				flags[c] = 1
				flags['.'] = 1
				break
			}
			lastc := chars[i+3]
			if c > lastc {
				// Invalid '..'-range, '..'-range needs to be incrementing
				flags[c] = 1
				flags[lastc] = 1
				flags['.'] = 1
			} else {
				for c_ := c; c_ < lastc; c_++ {
					flags[c_] = 1
				}
			}
		} else {
			flags[c] = 1
		}
	}
	// dst
	len_ = len(s)
	dst := make([]byte, len_)
	pos := 0

	for i := 0; i < len_; i++ {
		c := s[i]

		if pos == 0 && flags[c] != 0 {
			continue
		}
		dst[pos] = c
		pos++
	}

	return string(dst[0:pos])
}
