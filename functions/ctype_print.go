package functions

// Check for printable character(s)
func CtypePrint(s string) bool {
	s_len := len(s)
	if s_len <= 0 {
		return false
	}

	for i := 0; i < s_len; i++ {
		c := s[i]
		if c < 0x20 || 0x7E < c {
			return false
		}
	}
	return true
}
