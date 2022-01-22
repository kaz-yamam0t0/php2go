package functions

import (
	//"fmt"
	"regexp"
	"strings"
)

// Convert special characters to HTML entities
// 
// Actually the sypnosis of Htmlspecialchars() is like this:
// 
//  Htmlspecialchars(s string[, flags int[, encoding string[, double_encode bool]]]) 
// 
// [Note] You can also use `html.EscapeString()`
func Htmlspecialchars(s string, args ...interface{}) string {
	var flags int = ENT_COMPAT | ENT_SUBSTITUTE | ENT_HTML401
	var encoding string
	double_encode := true

	// args
	args_len := len(args)
	if args_len >= 1 {
		flags = args[0].(int)
	}
	if args_len >= 2 {
		encoding = args[1].(string)
	}
	if args_len >= 3 {
		double_encode = args[2].(bool)
	}

	// is `encoding` utf8?
	is_utf8 := false
	if encoding != "" {
		if regexp.MustCompile(`^utf\-?8$`).MatchString(strings.ToLower(encoding)) {
			is_utf8 = true
		}
	}

	// init dst
	dst := []byte("")
	bit_flags := [...]byte{0x80, 0xC0, 0xE0, 0xF0, 0xF8, 0xFE}

	s_len := len(s)

	// each char
	for i := 0; i < s_len; i++ {
		c := s[i]

		// culculate utf-8 size + validate chars
		invalid_ := false
		size_ := 1
		if is_utf8 {
			for j := 6; j > 1; j-- {
				if (c & bit_flags[j-1]) == bit_flags[j-2] {
					size_ = j

					if i+size_-1 >= s_len {
						size_ = s_len - i
						invalid_ = true
						break
					}

					j--
					for ; j > 1; j-- {
						if (s[i+(size_-j)] & bit_flags[j-1]) != bit_flags[j-2] {
							size_ = size_ - j + 1
							invalid_ = true
							break
						}
					}
				}
			}
		}
		if size_ > 1 {
			if c >= 0x80 {
				invalid_ = true
			}
		}
		// invalid chars
		if invalid_ {
			if (flags & ENT_IGNORE) == ENT_IGNORE {
				// ignore
			} else if (flags & ENT_SUBSTITUTE) == ENT_SUBSTITUTE {
				//size_--
				for ; size_ > 1; size_-- {
					dst = append(dst, 0xEF, 0xBF, 0xBD)
				}
			} else {
				return ""
			}
			continue
		}
		// more than 2 bytes characters will not be converted
		if size_ > 1 {
			dst = append(dst, c)

			for ; size_ > 1; size_-- {
				i++
				dst = append(dst, s[i])
			}
		} else {
			// 1 byte char
			if c == '<' {
				dst = append(dst, "&lt;"...)
			} else if c == '>' {
				dst = append(dst, "&gt;"...)
			} else if c == '"' || c == '\'' {
				if (flags&ENT_QUOTES) == ENT_QUOTES || ((flags&ENT_COMPAT) == ENT_COMPAT && c != '\'') {
					if s[i] == '"' {
						dst = append(dst, "&quot;"...)
					} else if s[i] == '\'' {
						dst = append(dst, "&#039;"...)
					}
				} else {
					dst = append(dst, c)
				}
			} else if c == '&' {
				if double_encode {
					dst = append(dst, "&amp;"...)
				} else {
					ent_flg := false
					n_flg := false
					alph_flg := false
					sharp_flg := 0

					pos := i + 1
					for ; pos < s_len; pos++ {
						if pos == i+1 && s[pos] == '#' {
							sharp_flg = 1
							if pos+1 < s_len && s[pos+1] == 'x' {
								sharp_flg = 2
							}
							continue
						}
						if ('a' <= s[pos] && s[pos] <= 'z') || ('A' <= s[pos] && s[pos] <= 'Z') {
							alph_flg = true
						} else if '0' <= s[pos] && s[pos] <= '9' {
							n_flg = true
						} else if s[pos] == ';' {
							ent_flg = true
						} else {
							break
						}
					}
					encode_flg := true
					if ent_flg {
						if sharp_flg == 1 && n_flg && !alph_flg {
							encode_flg = false
						} else if sharp_flg == 2 && (n_flg || alph_flg) {
							hex_flg := true
							for i2 := i + 3; i2 < pos-1; i2++ {
								if !(('0' < s[i2] && s[i2] <= '9') || ('a' < s[i2] && s[i2] <= 'f') || ('A' < s[i2] && s[i2] <= 'F')) {
									hex_flg = false
									break
								}
							}
							if hex_flg {
								encode_flg = false
							}
						} else {
							// i+1 - pos-1 までの間がエンティティ
							// encode_flg = false

							ent_ := s[i+1 : pos-1]
							if (flags & ENT_HTML5) == ENT_HTML5 {
								if getEntitiesApos().MatchString(ent_) || getEntitiesHTML5().MatchString(ent_) {
									encode_flg = false
								}
							} else if (flags&ENT_XML1) == ENT_XML1 || (flags&ENT_XHTML) == ENT_XHTML {
								if getEntitiesApos().MatchString(ent_) || getEntitiesXHTML().MatchString(ent_) {
									encode_flg = false
								}
							} else {
								if getEntities401().MatchString(ent_) {
									encode_flg = false
								}
							}
						}
					}
					if encode_flg {
						dst = append(dst, "&amp;"...)
					} else {
						dst = append(dst, c)
					}
				}
			} else {
				dst = append(dst, c)
			}
		}
	}
	return string(dst)

	/*
		Light version (The second and subsequent arguments are ignored)
		Actually this is enough in most cases

		pos := 0
		s_len := len(s)
		for i:=0; i<s_len; i++ {
			if s[i] == '<' || s[i] == '>' {
				pos += 4
			} else if s[i] == '&' {
				pos += 5
			} else if s[i] == '"' || s[i] == '\'' {
				pos += 6
			} else {
				pos++
			}
		}
		if pos <= 0 {
			return ""
		}
		dst := []byte("")
		for i:=0; i<s_len; i++ {
			if s[i] == '<' {
				dst = append(dst, "&lt;"...)
			} else if s[i] == '>' {
				dst = append(dst, "&gt;"...)
			} else if s[i] == '&' {
				dst = append(dst, "&amp;"...)
			} else if s[i] == '"' {
				dst = append(dst, "&quot;"...)
			} else if s[i] == '\'' {
				dst = append(dst, "&#039;"...)
			} else {
				dst = append(dst, s[i])
			}
		}
		return string(dst)

	*/
}
