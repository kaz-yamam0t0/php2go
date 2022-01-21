/**
 * Golang equivalent to php `rawurlencode`
 *
 * URL-encode according to RFC 3986
 * @see https://www.php.net/manual/en/function.rawurlencode.php
 *
 * @param string s
 * @return string
 */
package functions



func Rawurlencode(s string) string {
	s_len := len(s)
	pos := 0
	for i := 0; i < s_len; i++ {
		c := s[i]
		if urlShouldEscape(c) {
			pos += 3
		} else {
			pos++
		}
	}
	dst_len := pos
	dst := make([]byte, dst_len)

	pos = 0
	for i := 0; i < s_len; i++ {
		c := s[i]
		if urlShouldEscape(c) {
			dst[pos], dst[pos+1], dst[pos+2] = '%', upperhex[c>>4], upperhex[c&0x0F]
			pos += 3
		} else {
			dst[pos] = c
			pos++
		}
	}
	return string(dst[0:pos])
}
