package functions

// Binary safe case-insensitive string comparison of the first n characters
func Strncasecmp(s1 string, s2 string, length int) int {
	s1_len := len(s1)
	s2_len := len(s2)
	len_ := s1_len
	if len_ > s2_len {
		len_ = s2_len
	}
	if len_ > length {
		len_ = length
	}

	for i := 0; i < len_; i++ {
		if !cmpLower(s1[i], s2[i]) {
			return int(lower(s1[i])) - int(lower(s2[i]))
		}
	}
	if len_ == length {
		return 0
	}

	if s1_len > s2_len {
		return 1
	} else if s1_len < s2_len {
		return -1
	} else {
		return 0
	}
}
