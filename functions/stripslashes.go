package functions

// Un-quotes a quoted string
func Stripslashes(s string) string {
	s_len := len(s)
	dst := make([]byte, s_len)

	pos := 0
	for i := 0; i < s_len; i++ {
		if s[i] == '\\' {
			if i+1 >= s_len || s[i+1] != '\\' {
				continue
			}
		}
		dst[pos] = s[i]
		pos++
	}
	return string(dst[0:pos])
}
