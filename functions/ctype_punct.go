package functions

// Check for any printable character which is not whitespace or an alphanumeric character
func CtypePunct(s string) bool {
	s_len := len(s)
	if s_len <= 0 {
		return false
	}

	for i := 0; i < s_len; i++ {
		c := s[i]
		if c < 0x20 || 0x7E < c {
			return false
		}
		if ('0' <= c && c <= '9') || ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z') {
			return false
		}
	}
	return true
}
