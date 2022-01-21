/**
 * Golang equivalent to php `chunk_split`
 *
 * Split a string into smaller chunks
 * @see https://www.php.net/manual/en/function.chunk-split.php
 *
 * @param string s
 * @param int length
 * @param string separator
 * @return string
 */
package functions

import "math"

func ChunkSplit(s string, args ...interface{}) string {
	length := 76
	separator := "\r\n"
	if len(args) >= 1 {
		length = args[0].(int)
		if length <= 0 {
			panic("length must be greater than 0")
		}
	}
	if len(args) >= 2 {
		separator = args[1].(string)
	}

	s_len := len(s)
	dst := make([]byte, s_len+(s_len/length+1)*len(separator))

	sep := []byte(separator)
	sep_len := len(sep)

	pos := 0
	for i := 0; i < s_len; i += length {
		_length := int(math.Min(float64(length), float64(s_len-i)))
		copy(dst[pos:pos+_length], []byte(s[i:i+_length]))
		pos += _length

		copy(dst[pos:pos+sep_len], sep)
		pos += sep_len
	}

	return string(dst[:pos])
}
