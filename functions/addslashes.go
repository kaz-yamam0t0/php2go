/**
 * php `addslashes` with Golang
 *
 * Quote string with slashes
 * @see https://www.php.net/manual/en/function.addslashes.php
 *
 * @param string s
 * @return string
 */
package functions

func Addslashes(s string) string {
	len := len(s)
	pos := 0
	for i := 0; i < len; i++ {
		if s[i] == '\'' || s[i] == '"' || s[i] == '\\' || s[i] == 0 {
			pos++
		}
		pos++
	}

	dst := make([]byte, pos)
	pos = 0
	for i := 0; i < len; i++ {
		if s[i] == '\'' || s[i] == '"' || s[i] == '\\' || s[i] == 0 {
			dst[pos] = '\\'
			pos++
		}
		dst[pos] = s[i]
		pos++
	}
	return string(dst[0:pos])
}
