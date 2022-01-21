/**
 * Golang equivalent to php `stripcslashes`
 *
 * Un-quote string quoted with addcslashes
 * @see https://www.php.net/manual/en/function.stripcslashes.php
 *
 * @param string s
 * @return string
 */
package functions

func Stripcslashes(s string) string {
	s_len := len(s)
	dst := make([]byte, s_len)

	pos := 0
	for i := 0; i < s_len; i++ {
		if s[i] != '\\' {
			dst[pos] = s[i]
			pos++
			continue
		}
		if i+1 >= s_len {
			break
		}
		i++

		switch s[i] {
		case 'n':
			dst[pos] = 0x0A
		case 't':
			dst[pos] = 0x09
		case 'r':
			dst[pos] = 0x0D
		case 'a':
			dst[pos] = 0x07
		case 'b':
			dst[pos] = 0x08
		case 'v':
			dst[pos] = 0x0B
		case 'f':
			dst[pos] = 0x0C
		case '\\':
			dst[pos] = 0x5C
		}
		if dst[pos] != 0 {
			pos++
			continue
		}
		if i+2 >= s_len {
			break
		}
		if is_digit(s[i]) && is_digit(s[i+1]) && is_digit(s[i+2]) {
			c1 := byte(s[i] - '0')
			c2 := byte(s[i+1] - '0')
			c3 := byte(s[i+2] - '0')
			dst[pos] = c1*64 + c2*8 + c3
			pos++
			i += 2
			continue
		}
		if s[i] == 'x' && is_hex(s[i+1]) && is_hex(s[i+2]) {
			c1 := c_hex(s[i+1])
			c2 := c_hex(s[i+2])
			dst[pos] = c1*16 + c2
			pos++
			i += 2
			continue
		}
		dst[pos] = s[i]
		pos++
	}
	return string(dst[0:pos])
}
