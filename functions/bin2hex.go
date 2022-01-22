package functions

// Convert binary data into hexadecimal representation
func Bin2hex(s string) string {
	s_len := len(s)

	hex_ := "0123456789abcdef"
	dst := make([]byte, s_len*2)

	for i := 0; i < s_len; i++ {
		c := s[i]
		dst[i*2] = hex_[(c >> 4)]
		dst[i*2+1] = hex_[c&0x0F]
	}
	return string(dst)
}
