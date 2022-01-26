package functions

// Strip whitespace (or other characters) from the end of a string
//
// Actually the sypnosis of Rtrim() is like this:
//
//  Rtrim(s string[, characters sring = " \n\r\t\v\x00"])
func Rtrim(s string, args ...string) string {
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

	i2 := -1

	for i := 0; i < len_; i++ {
		c := s[i]

		if flags[c] != 0 {
			if i2 < 0 {
				i2 = i
			}
		} else {
			if i2 >= 0 {
				for j := i2; j < i; j++ {
					dst[pos] = s[j]
					pos++
				}
				i2 = 0
			}
			dst[pos] = c
			pos++
			i2 = -1
		}
	}

	return string(dst[0:pos])
}
