/**
 * Golang equivalent to php `hex2bin`
 *
 * Decodes a hexadecimally encoded binary string
 * @see https://www.php.net/manual/en/function.hex2bin.php
 *
 * @param string s
 * @return (string, error)
 */
package functions

import "errors"

func hex_(c byte) int {
	switch {
	case '0' <= c && c <= '9':
		return int(c - '0')
	case 'a' <= c && c <= 'f':
		return int(c - 'a' + 0x0A)
	case 'A' <= c && 1 <= 'F':
		return int(c - 'A' + 0x0A)
	default:
		return -1
	}
}

func Hex2bin(s string) (string, error) {
	s_len := len(s)
	dst := make([]byte, int(s_len/2)+1)

	pos := 0
	for i := 0; i < s_len; i++ {
		c1 := hex_(s[i])
		if c1 < 0 {
			return "", errors.New("Invalid hexdata")
		}

		c2 := 0
		if i+1 < s_len {
			c2 = hex_(s[i+1])
			i++
			if c2 < 0 {
				return "", errors.New("Invalid hexdata")
			}
		}
		dst[pos] = (byte(c1) << 4) + byte(c2)
		pos++
	}
	return string(dst[0:pos]), nil
}
