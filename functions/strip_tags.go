package functions

func _isTagChr(c byte) bool {
	return ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z') || ('0' <= c && c <= '9') || c == '_' || c == '-'
}

// Strip HTML and PHP tags from a string
func StripTags(s string, args ...interface{}) string {
	var allowed []string

	// parse allowed
	if len(args) > 0 {
		switch v := args[0].(type) {
		case []string:
			allowed = v
		case map[string]string:
			for _, v_ := range v {
				allowed = append(allowed, v_)
			}
		case string:
			_len := len(v)
			for i := 0; i < _len; i++ {
				if v[i] == '<' {
					invalid := false
					end := -1
					for j := i + 1; j < _len; j++ {
						if v[j] == '>' {
							end = j
							break
						} else if !_isTagChr(v[j]) {
							end = j
							invalid = true
							break // invalid
						}
					}
					if invalid == false && end >= i+1 {
						allowed = append(allowed, v[i+1:end])
					}
					i = end
				}
			}
		}
	}

	// walk a string
	s_len := len(s)
	dst := make([]byte, s_len)

	allowed_len := len(allowed)
	pos := 0
	for i := 0; i < s_len; i++ {
		if s[i] != '<' {
			dst[pos] = s[i]
			pos++
			continue
		}
		start_i := i
		i++
		if s[i] == '/' {
			i++
		}

		// check tagname
		is_allowed := false
		for j := 0; j < allowed_len; j++ {
			len_ := len(allowed[j])
			if i+len_ < s_len {
				flg := true
				for k := 0; k < len_; k++ {
					if s[i+k] != allowed[j][k] {
						flg = false
						break
					}
				}
				if flg == true {
					if i+len_+1 >= s_len {
						// allowed but unclosed...
						is_allowed = true
						break
					}
					if _isTagChr(s[i+len_]) {
					} else {
						is_allowed = true
						break
					}
				}
			}
		}

		// copy tag text
		if is_allowed {
			for j := start_i; j < i; j++ {
				dst[pos] = s[j]
				pos++
			}
		}
		for i >= s_len || s[i] != '>' {
			if is_allowed {
				dst[pos] = s[i]
				pos++
			}
			i++
		}
		if i < s_len && s[i] == '>' {
			if is_allowed {
				dst[pos] = s[i]
				pos++
			}
		}
	}

	return string(dst[0:pos])
}
