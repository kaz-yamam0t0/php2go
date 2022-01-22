package functions

// Check for control character(s)
func CtypeCntrl(s string) bool {
	s_len := len(s)
	if s_len <= 0 {
		return false
	}

	for i := 0; i < s_len; i++ {
		c := s[i]
		if c >= 0x20 {
			return false
		}
	}
	return true
}
