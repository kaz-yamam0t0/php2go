package functions

//  Quote string with slashes in a C style
func Addcslashes(s string, chars string) string {
	var flags [256]byte

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
	// culc dst length
	len_ = len(s)
	dst_len := 0
	for i := 0; i < len_; i++ {
		c := s[i]
		if flags[c] != 0 {
			if c < 32 || 126 < c {
				switch c {
				case 0x0a, 0x09, 0x0d, 0x07, 0x08, 0x0b, 0x0c:
					dst_len += 2
				default:
					dst_len += 4
				}
				continue
			}
			dst_len++
		}
		dst_len++
	}

	dst := make([]byte, dst_len)
	pos := 0

	for i := 0; i < len_; i++ {
		c := s[i]

		if flags[c] != 0 {
			if c < 32 || 126 < c {
				dst[pos] = '\\'
				pos++
				switch c {
				case 0x0a:
					dst[pos] = 'n'
				case 0x09:
					dst[pos] = 't'
				case 0x0d:
					dst[pos] = 'r'
				case 0x07:
					dst[pos] = 'a'
				case 0x08:
					dst[pos] = 'b'
				case 0x0b:
					dst[pos] = 'v'
				case 0x0c:
					dst[pos] = 'f'
				default:
					dst[pos] = '0' + (byte(c/64) % 8)
					dst[pos+1] = '0' + (byte(c/8) % 8)
					dst[pos+2] = '0' + (byte(c) % 8)
					pos += 2
				}
				pos++
				continue
			}
			dst[pos] = '\\'
			pos++
		}
		dst[pos] = c
		pos++
	}

	return string(dst)
}
