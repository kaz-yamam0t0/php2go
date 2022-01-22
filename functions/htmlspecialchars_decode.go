package functions

// Convert special HTML entities back to characters
// 
// Actually the sypnosis of HtmlspecialcharsDecode() is like this:
// 
//  HtmlspecialcharsDecode(s string[, flags int]) 
func HtmlspecialcharsDecode(s string, args ...int) string {
	var flags int = ENT_COMPAT | ENT_SUBSTITUTE | ENT_HTML401
	if len(args) >= 1 {
		flags = args[0]
	}

	s_len := len(s)
	dst := make([]byte, s_len)

	pos := 0
	for i := 0; i < s_len; i++ {
		c := s[i]
		if c == '&' && s_len-i >= 4 {
			switch {
			case s_len-i >= 6 && s[i:i+6] == "&quot;":
				if (flags & ENT_COMPAT) == ENT_COMPAT {
					dst[pos] = '"'
					pos++
					i += 5
					continue
				}
			case s_len-i >= 6 && s[i:i+6] == "&#039;":
				if (flags & ENT_QUOTES) == ENT_QUOTES {
					dst[pos] = '\''
					pos++
					i += 5
					continue
				}
			case s_len-i >= 5 && s[i:i+5] == "&amp;":
				dst[pos] = '&'
				pos++
				i += 4
				continue
			case s_len-i >= 4 && s[i:i+4] == "&lt;":
				dst[pos] = '<'
				pos++
				i += 3
				continue
			case s_len-i >= 4 && s[i:i+4] == "&gt;":
				dst[pos] = '>'
				pos++
				i += 3
				continue
			}
		}
		dst[pos] = c
		pos++
	}
	return string(dst[0:pos])
}
