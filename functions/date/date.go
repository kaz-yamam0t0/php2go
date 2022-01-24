package date

import (
	"time"
)

// Format a local time/date
// 
// Actually the sypnosis of Date() is like this:
// 
//  Date(s string, baseTime int*|time.Time)
func DateFormat(s string, dt time.Time) string {
	var dst []byte
	s_len := len(s)
	pos := 0
	for i := 0; i < s_len; i++ {
		// backslash
		if s[i] == '\\' {
			dst = append(dst, s[pos:i]...)
			if i+1 < s_len {
				dst = append(dst, s[i+1])
			}
			pos = i + 2
			i++ // skip
			continue
		}
		if b, ok := DateFormatChar(s[i], dt); ok == true {
			if pos < i {
				dst = append(dst, s[pos:i]...)
			}
			dst = append(dst, b...)
			pos = i + 1
		}
	}
	if pos < s_len {
		dst = append(dst, s[pos:]...)
	}
	return string(dst)
}
