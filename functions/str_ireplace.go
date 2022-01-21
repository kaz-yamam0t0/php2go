/**
 * Golang equivalent to php `str_ireplace`
 *
 * Case-insensitive version of str_replace
 * @see https://www.php.net/manual/en/function.str-ireplace.php
 *
 * @param string search
 * @param string replace
 * @param string subject
 * @return string
 */
package functions

func StrIreplace(search string, replace string, s string) string {
	s_len := len(s)
	search_len := len(search)
	replace_len := len(replace)

	// culculate length
	pos := 0
	for i := 0; i < s_len; i++ {
		match := false
		if s_len-i >= search_len {
			match = true
			for j := 0; j < search_len; j++ {
				if !cmpLower(s[i+j], search[j]) {
					match = false
					break
				}
			}
		}
		if match {
			i += (search_len - 1)
			pos += replace_len
		} else {
			pos++
		}
	}

	// copy
	dst := make([]byte, pos)

	pos = 0
	for i := 0; i < s_len; i++ {
		match := false
		if s_len-i >= search_len {
			match = true
			for j := 0; j < search_len; j++ {
				if !cmpLower(s[i+j], search[j]) {
					match = false
					break
				}
			}
		}
		if match {
			for j := 0; j < replace_len; j++ {
				dst[pos] = replace[j]
				pos++
			}
			i += (search_len - 1)
		} else {
			dst[pos] = s[i]
			pos++
		}
	}

	return string(dst)
}
