package date

// ==============================================================
// get*Table
// ==============================================================

// get month table
func getMonthTable() map[string]int {
	return map[string]int {
		"january": 1, 
		"jan": 1, 
		"february": 2,
		"feb": 2,
		"march" : 3,
		"mar": 3,
		"april":4,
		"apr":4,
		"may":5,
		"june":6,
		"jun":6,
		"july":7,
		"jul":7,
		"august":8,
		"aug":8,
		"september":9,
		"sep":9,
		"octobar": 10,
		"oct":10,
		"november": 11,
		"nov":11,
		"december": 12,
		"dec":12, 
	}
}
// get weekday table 
func getWeekdayTable() map[string]int {
	return map[string]int {
		"sunday": 0,
		"monday": 1,
		"tuesday": 2,
		"wednesday": 3,
		"thursday": 4,
		"friday": 5,
		"saturday": 6,
	}
}

// ==============================================================
// is*
// ==============================================================

func isSpace(c byte) bool {
	return c == 0 || c == 0x20 || c == '\n' || c == '\r' || c == '\t' || c == '\v'
}
func isSeparator(c byte) bool {
	return c == ';' || c == ':' || c == '/' || c == '.' || c == ',' || c == '-'  || c == '(' || c == ')'
}
func isNumeric(c byte) bool {
	return '0' <= c && c <= '9'
}
func isAlphanumeric(c byte) bool {
	return ('0' <= c && c <= '9') || ('a' <= c && c <= 'z') || ('A' <= c && c <= 'Z')
}

// ==============================================================
// skip*
// ==============================================================

func skipSpaces(s *string, pos_s *int) int {
	return skipChars(s, pos_s, isSpace)
}
func skipSeparators(s *string, pos_s *int) int {
	return skipChars(s, pos_s, isSeparator)
}

func skipChars(s *string, pos_s *int, callback func(c byte) bool) int {
	start_pos := *pos_s
	s_len := len(*s)
	for *pos_s < s_len {
		if callback((*s)[*pos_s]) {
			(*pos_s)++
		} else {
			break
		}
	}
	return (*pos_s) - start_pos
}

