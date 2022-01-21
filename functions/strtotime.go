/**
 * Golang equivalent to php `strtotime`
 *
 * Parse about any English textual datetime description into a Unix timestamp
 * @see https://www.php.net/manual/en/function.strtotime.php
 *
 * @param string datetime
 * @param int baseTimestamp
 * @return interface{}
 */
package functions

import (
	"regexp"
	"strconv"
	"strings"
	"time"
)

// cache regexp.Regexp
var p_ignore *regexp.Regexp
var p_isoformat *regexp.Regexp
var p_norm_sp *regexp.Regexp

var p_sp *regexp.Regexp
var p_now *regexp.Regexp
var p_dmy *regexp.Regexp
var p_ymd *regexp.Regexp
var p_time *regexp.Regexp
var p_tz *regexp.Regexp
var p_months *regexp.Regexp
var p_weeks *regexp.Regexp
var p_relative_pos *regexp.Regexp
var p_pos *regexp.Regexp

// should be called before each regexp variables are being used
var init_flg bool

func initRegexp() {
	if init_flg {
		return
	}
	init_flg = true

	_months := "(january|february|march|april|may|june|july|august|september|october|november|december|jan|feb|mar|apr|jun|jul|aug|sep|oct|nov|dec)"
	_weeks := "(sunday|monday|tuesday|wednesday|thursday|friday|saturday|sun|mon|tue|wed|thu|fri|sat)"

	p_ignore = regexp.MustCompile(`(\bthe\b|,)`)
	p_isoformat = regexp.MustCompile(`\b(\d{4}\-\d{2}\-\d{2})t(\d{2}:\d{2}(:\d{2})?)([\-\+]?\d{2}\:?\d{2})?\b`)
	p_norm_sp = regexp.MustCompile(`\s+`)

	//p_sp        = regexp.MustCompile(`^[\s\n\r\t]$`)
	p_now = regexp.MustCompile(`^now\b`)                                                                                                                                        // Now
	p_dmy = regexp.MustCompile(`^(\d{1,2})(st|nd|rd|th)?\s*` + _months + `(\s+(\d+))?\b`)                                                                                       // 1st January 2006
	p_ymd = regexp.MustCompile(`^(\d+)([/\-\.])(\d+)([/\-\.])(\d+)\b`)                                                                                                          // 2006-01-01
	p_time = regexp.MustCompile(`^(\d{2})\:(\d{2})(\:(\d{2}))?( (a\.m\.|p\.m\.|am|pm))?\b`)                                                                                     // 00:00:00 am
	p_tz = regexp.MustCompile(`^([\-\+]?)(\d{2})\:?(\d{2})\b`)                                                                                                                  // +09:00
	p_months = regexp.MustCompile(`^` + _months + `\b`)                                                                                                                         // Junuary
	p_weeks = regexp.MustCompile(`^` + _weeks + `\b`)                                                                                                                           // Sunday
	p_relative_pos = regexp.MustCompile(`^([\+\-]?)\s*(\d+|a)\s*(year|month|day|hour|minute|second|week|millisecond|microsecond|msec|ms|µsec|µs|usec|sec|min|forth?night)s?\b`) // +1 year
	p_pos = regexp.MustCompile(`^((first|last) day of )?(next|last) ((year|month|day)|` + _months + `|` + _weeks + `)\b`)                                                       // +1 year
}

/*
// not used.

func freeRegexp() {
	//if init_flg == false {
	//	return
	//}
	init_flg = false

	p_ignore    = nil
	p_isoformat = nil
	p_norm_sp   = nil

	p_sp        = nil
	p_now       = nil
	p_dmy       = nil
	p_ymd       = nil
	p_time      = nil
	p_tz        = nil
	p_months    = nil
	p_weeks     = nil
	p_relative_pos = nil
	p_pos = nil


}  */

// Relative Additions or Subtractions
type timeAddition struct {
	n    int
	unit string

	pos     string
	month   int
	weekday int
	word    string
	day_flg string
}

// Time Data
type timeData struct {
	y         int
	m         int
	d         int
	h         int
	i         int
	s         int
	us        int
	day       int // actually doesn't affect the result
	z         int
	additions []timeAddition
}

// check if ymd is correct
func checkDate(y int, m int, d int) bool {
	if y < 0 || (m < 1 || 12 < m) || (d < 1 || 31 < d) {
		return false
	}
	return d <= getLastDay(y, m)
}

// get the last day of a month
func getLastDay(y int, m int) int {
	if m == 4 || m == 6 || m == 9 || m == 11 {
		return 30
	}
	if m == 2 {
		if y%4 == 0 && (y%100 != 0 || y%400 == 0) {
			return 29
		} else {
			return 28
		}
	}
	return 31
}

// month name string to number
func getMonthNum(month_name string) int {
	// month_name should be the lower cases of a correct month name
	switch {
	case month_name == "january" || month_name == "jan":
		return 1
	case month_name == "february" || month_name == "feb":
		return 2
	case month_name == "march" || month_name == "mar":
		return 3
	case month_name == "april" || month_name == "apr":
		return 4
	case month_name == "may":
		return 5
	case month_name == "june" || month_name == "jun":
		return 6
	case month_name == "july" || month_name == "jul":
		return 7
	case month_name == "august" || month_name == "aug":
		return 8
	case month_name == "september" || month_name == "sep":
		return 9
	case month_name == "october" || month_name == "oct":
		return 10
	case month_name == "november" || month_name == "nov":
		return 11
	case month_name == "december" || month_name == "dec":
		return 12
	}
	return -1
}

// weekday name string to number
func getWeekdayNum(day_name string) int {
	// day_name should be the lower cases of a correct day name
	switch {
	case day_name == "sunday" || day_name == "sun":
		return 0
	case day_name == "monday" || day_name == "mon":
		return 1
	case day_name == "tuesday" || day_name == "tue":
		return 2
	case day_name == "wednesday" || day_name == "wed":
		return 3
	case day_name == "thursday" || day_name == "thu":
		return 4
	case day_name == "friday" || day_name == "fri":
		return 5
	case day_name == "saturday" || day_name == "sat":
		return 6
	}
	return -1
}

// scan format chunk and add pos
func scanFormat(data *timeData, s string, pos int) int {
	s_len := len(s)
	if pos >= s_len {
		return -1
	}
	// \s\n\r\t
	for s_len > pos && (s[pos] == 0x20 || s[pos] == '\n' || s[pos] == '\r' || s[pos] == '\t') {
		pos++
	}
	if pos >= s_len {
		return pos
	}

	_s := s[pos:]
	_s_len := len(_s)
	_ = _s_len

	if p_now.MatchString(_s) {
		// Now ===============
		d_ := time.Now()

		data.y = d_.Year()
		data.m = int(d_.Month())
		data.d = d_.Day()
		data.h = d_.Hour()
		data.i = d_.Minute()
		data.s = d_.Second()
		data.us = int(d_.Nanosecond() / 1e3)
		pos += 3
	} else if _g := p_months.FindStringSubmatch(_s); len(_g) > 0 {
		// January ===============
		data.m = getMonthNum(_g[1])
		pos += len(_g[0])
	} else if _g := p_weeks.FindStringSubmatch(_s); len(_g) > 0 {
		// Sunday ================
		data.day = getWeekdayNum(_g[1])
		pos += len(_g[0])
	} else if _g := p_relative_pos.FindStringSubmatch(_s); len(_g) > 0 {
		// 1 year,month,day,hour,minute,week,millisecond,microseconds =========
		sign_ := _g[1]
		n_ := 1
		unit_ := _g[3]

		if tmp, err := strconv.Atoi(_g[2]); err == nil {
			n_ = tmp
		}
		if sign_ == "-" {
			n_ = n_ * -1
		}

		data.additions = append(data.additions, timeAddition{
			n:    n_,
			unit: unit_,
		})
		pos += len(_g[0])
	} else if _g := p_pos.FindStringSubmatch(_s); len(_g) > 0 {
		// (first|last day of) first|next year|month|day
		day_flg_ := _g[2]
		pos_flg_ := _g[3]
		word_ := _g[5]
		month_name_ := _g[6]
		day_name_ := _g[7]

		a := timeAddition{
			day_flg: day_flg_, // empty | first | last (day of)
			pos:     pos_flg_, // first | next
			month:   -1,
			weekday: -1,
			word:    "",
		}
		if month_name_ != "" {
			a.month = getMonthNum(month_name_)
		} else if day_name_ != "" {
			a.weekday = getWeekdayNum(day_name_)
		} else {
			a.word = word_
		}

		data.additions = append(data.additions, a)
		pos += len(_g[0])
	} else if _g := p_dmy.FindStringSubmatch(_s); len(_g) > 0 {
		// 1st January 2021 ============
		data.d, _ = strconv.Atoi(_g[1])
		data.m = getMonthNum(_g[3])
		data.y, _ = strconv.Atoi(_g[5])
		pos += len(_g[0])
	} else if _g := p_ymd.FindStringSubmatch(_s); len(_g) > 0 {
		// y/m/d y-m-d m/d/y d.m.y ==============
		if _g[2] != _g[4] {
			return -1
		}
		sep_ := _g[2]
		d1_, _ := strconv.Atoi(_g[1])
		d2_, _ := strconv.Atoi(_g[3])
		d3_, _ := strconv.Atoi(_g[5])

		switch sep_[0] {
		case '.': //  d.m.y
			if !checkDate(d3_, d2_, d1_) {
				return -1
			}
			data.y, data.m, data.d = d3_, d2_, d1_
		case '-': //  y-m-d d-m-y
			if checkDate(d3_, d2_, d1_) {
				data.y, data.m, data.d = d3_, d2_, d1_
			} else if checkDate(d1_, d2_, d3_) {
				data.y, data.m, data.d = d1_, d2_, d3_
			} else {
				return -1
			}
		case '/': // y/m/d m/d/y
			if checkDate(d3_, d1_, d2_) {
				data.y, data.m, data.d = d3_, d1_, d2_
			} else if checkDate(d1_, d2_, d3_) {
				data.y, data.m, data.d = d1_, d2_, d3_
			} else {
				return -1
			}
		}
		pos += len(_g[0])
	} else if _g := p_time.FindStringSubmatch(_s); len(_g) > 0 {
		// 00:00:00 am ============
		_h, _ := strconv.Atoi(_g[1])
		_i, _ := strconv.Atoi(_g[2])
		_pm := strings.Replace(_g[6], ".", "", -1)
		_s := -1

		if _g[4] != "" {
			_s, _ = strconv.Atoi(_g[4])
		}
		if (_h < 0 || 24 <= _h) || (_i < 0 || 60 <= _i) || 60 <= _s {
			return -1
		}
		if _pm == "am" {
			if _h > 12 {
				return -1
			}
		} else if _pm == "pm" {
			if _h != 12 {
				_h += 12
			}
		}
		data.h = _h
		data.i = _i
		if _s >= 0 {
			data.s = _s
		}
		pos += len(_g[0])
	} else if _g := p_tz.FindStringSubmatch(_s); len(_g) > 0 {
		// +00:00 ======================
		_sign := _g[1]
		_h, _ := strconv.Atoi(_g[2])
		_i, _ := strconv.Atoi(_g[3])

		data.z = (_h*60 + _i) * 60
		if _sign == "-" {
			data.z = data.z * -1
		}
	} else {
		return -1
	}

	return pos
}

func move(data *timeData, a *timeAddition) {
	if a.unit != "" {
		add(data, a)
		return
	}
	switch {
	case a.month > 0:
		if a.pos == "next" && data.m >= a.month {
			data.y++
		} else if a.pos == "last" && data.m <= a.month {
			data.y--
		}
		data.m = a.month
	case a.weekday >= 0:
		w_ := int(time.Date(data.y, time.Month(data.m), data.d, data.h, data.i, data.s, data.us*1e3, time.Local).Weekday())
		if a.pos == "next" {
			n := (a.weekday+7-1-w_)%7 + 1
			a.n, a.unit = n, "day"
			add(data, a)
		} else if a.pos == "last" {
			n := -((w_+7-1-a.weekday)%7 + 1)
			a.n, a.unit = n, "day"
			add(data, a)
		}
	case a.word != "":
		if a.word != "year" && a.word != "month" && a.word != "day" {
			panic("unknown word") // never be occurred
		}
		n := -1
		if a.pos == "next" {
			n = 1
		}
		a.n, a.unit = n, a.word
		add(data, a)
	}

}
func add(data *timeData, a *timeAddition) {
	// might be outside of the range
	// it will be normalized when data is instantiated.
	switch a.unit {
	case "year":
		data.y += a.n
	case "month":
		data.m += a.n
	case "day":
		data.d += a.n
	case "hour":
		data.h += a.n
	case "min":
		fallthrough
	case "minute":
		data.i += a.n
	case "sec":
		fallthrough
	case "second":
		data.s += a.n
	case "ms":
		fallthrough
	case "msec":
		fallthrough
	case "millisecond":
		data.us += a.n * 1e3
	case "µs":
		fallthrough
	case "µsec":
		fallthrough
	case "microsecond":
		data.us += a.n
	case "week":
		data.d += a.n * 7
	case "forthnight":
		fallthrough
	case "fortnight":
		data.d += a.n * 14
	default:
		panic("unsupported units")
	}
	normalize(data)
}
func normalize(data *timeData) {
	data.s, data.us = norm(data.s, data.us, 1000000)
	data.i, data.s = norm(data.i, data.s, 60)
	data.h, data.i = norm(data.h, data.i, 60)
	data.d, data.h = norm(data.d, data.h, 24)

	// day, month, year
	normalizeYmd(data)
}

func normalizeYmd(data *timeData) {
	// year, month
	m := data.m - 1
	data.y, m = norm(data.y, m, 12)
	data.m = m + 1

	// days
	for data.d <= 0 {
		// last month
		data.m--
		if data.m <= 0 {
			data.m = 12
			data.y--
		}
		// days of the last month
		data.d += getLastDay(data.y, data.m)
	}

	last_d := getLastDay(data.y, data.m)
	for data.d > last_d {
		data.m++
		if data.m > 12 {
			data.m = 1
			data.y++
		}
		data.d -= last_d

		last_d = getLastDay(data.y, data.m)
	}
}

// copied from time/time.go
// @see https://cs.opensource.google/go/go/+/refs/tags/go1.17.5:src/time/time.go;l=1352;drc=refs%2Ftags%2Fgo1.17.5
func norm(hi, lo, base int) (nhi, nlo int) {
	if lo < 0 {
		n := (-lo-1)/base + 1
		hi -= n
		lo += n * base
	}
	if lo >= base {
		n := lo / base
		hi += n
		lo -= n * base
	}
	return hi, lo
}

func parseTimeFormat(format string, base *time.Time) (*time.Time, *timeData) {
	s := strings.TrimSpace(strings.ToLower(format))
	if s == "" {
		return nil, nil
	}

	initRegexp()
	//defer freeRegexp()

	// data
	data := timeData{-1, -1, -1, 0, 0, 0, 0, 0, 0, []timeAddition{}}

	if base == nil {
		t_ := time.Now()
		base = &t_
	}

	data.y = base.Year()
	data.m = int(base.Month())
	data.d = base.Day()
	data.h = base.Hour()
	data.i = base.Minute()
	data.s = base.Second()
	data.us = int(base.Nanosecond() / 1e3)

	_, data.z = base.Zone()

	// convert datetime
	s = p_ignore.ReplaceAllString(s, ` `)
	s = p_isoformat.ReplaceAllString(s, `$1 $2 $4`)
	s = p_norm_sp.ReplaceAllString(s, ` `)

	// parse string
	s_len := len(s)
	pos := 0
	cnts := 0
	for cnts < 15 && pos < s_len {
		pos = scanFormat(&data, s, pos)
		if pos < 0 {
			return nil, nil
		}
		cnts++
	}
	//fmt.Print("[data]",data,"\n")

	// additions
	a_len := len(data.additions)
	for i := 0; i < a_len; i++ {
		move(&data, &data.additions[i])
	}
	//fmt.Print("[data] -> ",data,"\n")

	// result
	res := time.Date(data.y, time.Month(data.m), data.d, data.h, data.i, data.s, data.us*1e3, time.Local)
	return &res, &data
}
func Str2Time(format string, base *time.Time) *time.Time {
	res, _ := parseTimeFormat(format, base)
	return res
}

func Strtotime(format string, args ...interface{}) int64 {
	// parse base
	var base time.Time
	if len(args) > 0 {
		switch tmp := args[0].(type) {
		case int:
			base = time.Unix(int64(tmp), 0)
		case int8:
			base = time.Unix(int64(tmp), 0)
		case int16:
			base = time.Unix(int64(tmp), 0)
		case int32:
			base = time.Unix(int64(tmp), 0)
		case int64:
			base = time.Unix(tmp, 0)
		case time.Time:
			base = tmp
		default:
			panic("unsupported argument")
		}
	} else {
		t_ := time.Now()
		base = time.Date(t_.Year(), t_.Month(), t_.Day(), 0, 0, 0, 0, t_.Location())
	}
	res, data := parseTimeFormat(format, &base)
	if res == nil {
		return -1
	}

	// convert to unix time
	t := res.Unix()
	_, offset := res.Zone()
	t += int64(offset)
	t -= int64(data.z)

	return t
}
