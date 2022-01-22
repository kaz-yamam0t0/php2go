package functions

// Inserts HTML line breaks before all newlines in a string
func Nl2br(s string, args ...bool) string {
	use_xhtml := true
	if len(args) > 0 {
		use_xhtml = args[0]
	}
	// br
	var br string
	if use_xhtml {
		br = "<br />"
	} else {
		br = "<br>"
	}
	br_len := len(br) // <br />

	// culculate length
	pos := 0
	s_len := len(s)
	for i := 0; i < s_len; i++ {
		if s[i] == 0x0A || s[i] == 0x0D {
			pos += br_len
		}
		pos++

		if i+1 < s_len && ((s[i] == 0x0A && s[i+1] == 0x0D) || (s[i] == 0x0D && s[i+1] == 0x0A)) {
			i++
			pos++
		}
	}
	// cpy
	dst := make([]byte, pos)
	pos = 0
	for i := 0; i < s_len; i++ {
		if s[i] == 0x0A || s[i] == 0x0D {
			for j := 0; j < br_len; j++ {
				dst[pos] = br[j]
				pos++
			}
		}
		dst[pos] = s[i]
		pos++

		if i+1 < s_len && ((s[i] == 0x0A && s[i+1] == 0x0D) || (s[i] == 0x0D && s[i+1] == 0x0A)) {
			i++
			dst[pos] = s[i]
			pos++
		}
	}

	return string(dst)
}
