package date

import (
	"regexp"
	"strconv"
	"strings"
	"time"
	//"fmt"
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
var p_iso *regexp.Regexp

// should be called before each regexp variables are being used
var date_regexp_init_flg bool

func initRegexp() {
	if date_regexp_init_flg {
		return
	}
	date_regexp_init_flg = true

	_months := "(january|february|march|april|may|june|july|august|september|october|november|december|jan|feb|mar|apr|jun|jul|aug|sep|oct|nov|dec)"
	_weeks := "(sunday|monday|tuesday|wednesday|thursday|friday|saturday|sun|mon|tue|wed|thu|fri|sat)"

	p_ignore = regexp.MustCompile(`(\bthe\b|,)`)
	p_isoformat = regexp.MustCompile(`\b(\d{4}\-\d{2}\-\d{2})t(\d{2}:\d{2}(:\d{2})?)([\-\+]?\d{2}\:?\d{2})?\b`)
	p_norm_sp = regexp.MustCompile(`\s+`)

	//p_sp        = regexp.MustCompile(`^[\s\n\r\t]$`)
	p_now = regexp.MustCompile(`^now\b`)   // Now
	p_dmy = regexp.MustCompile(`^(\d{1,2})(st|nd|rd|th)?\s*` + _months + `(\s+(\d+))?\b`)                                                                                       // 1st January 2006
	p_ymd = regexp.MustCompile(`^(\d+)([/\-\.])(\d+)([/\-\.])(\d+)\b`)                                                                                                          // 2006-01-01
	p_time = regexp.MustCompile(`^(\d{2})\:(\d{2})(\:(\d{2}))?( (a\.m\.|p\.m\.|am|pm))?\b`)                                                                                     // 00:00:00 am
	p_tz = regexp.MustCompile(`^([\-\+]?)(\d{2})\:?(\d{2})\b`)                                                                                                                  // +09:00
	p_months = regexp.MustCompile(`^` + _months + `\b`)                                                                                                                         // Junuary
	p_weeks = regexp.MustCompile(`^` + _weeks + `\b`)                                                                                                                           // Sunday
	p_relative_pos = regexp.MustCompile(`^([\+\-]?)\s*(\d+|a)\s*(year|month|day|hour|minute|second|week|millisecond|microsecond|msec|ms|µsec|µs|usec|sec|min|forth?night)s?(\s+ago)?\b`) // +1 year
	p_pos = regexp.MustCompile(`^((first|last) day of )?(next|last) ((year|month|day)|` + _months + `|` + _weeks + `)\b`)                                                       // +1 year
	p_iso = regexp.MustCompile(`^p((\d+)y)?((\d+)m)?((\d+)d)?(t((\d+)h)?((\d+)m)?((\d+(\.\d+)?)s)?)?\b`)                                     // P1Y2M3DT4H5M6.123456S
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

	p_iso = nil

}  */


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
func getMonthNum(s string) int {
	// month_name should be the lower cases of a correct month name
	for month_name, month_num := range getMonthTable() {
		if s == month_name {
			return month_num
		}
	}
	return -1
}
// starts with month name
func startsWithMonthName(s string) (int, string) {
	// month_name should be the lower cases of a correct month name
	for month_name, month_num := range getMonthTable() {
		if strings.HasPrefix(s, month_name) {
			return month_num, month_name
		}
	}
	return -1, ""
} 


// weekday name string to number
func getWeekdayNum(s string) int {
	// day_name should be the lower cases of a correct day name
	for weekday_name, weekday_num := range getWeekdayTable() {
		if s == weekday_name {
			return weekday_num
		}
	}
	return -1
}
// starts with month name
func startsWithWeekdayName(s string) (int, string) {
	// month_name should be the lower cases of a correct month name
	flg3 := len(s) > 3
	for weekday_name, weekday_num := range getWeekdayTable() {
		if strings.HasPrefix(s, weekday_name) || (flg3 && s[:3] == weekday_name[:3]) {
			return weekday_num, weekday_name
		}
	}
	return -1, ""
} 



// scan format chunk and add pos
func scanFormat(data *TimeData, s string, pos int) int {
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
		ago_ := _g[4]

		if tmp, err := strconv.Atoi(_g[2]); err == nil {
			n_ = tmp
		}
		if sign_ == "-" {
			n_ = n_ * -1
		}
		if ago_ != "" {
			n_ = n_ * -1
		}

		data.additions = append(data.additions, TimeAddition{
			n:    n_,
			unit: unit_,
		})
		pos += len(_g[0])
	} else if _g := p_iso.FindStringSubmatch(_s); len(_g) > 0 {
		// P1Y2M3DT4H5M6.7S =========
		//for unit_, index_ := range map[string]int{"year":2,"month":4,"day":6,"hour":9,"minute":11,"sec":13} 
		for unit_, index_ := range map[string]int{"year":2,"month":4,"day":6,"hour":9,"minute":11} {
			if _g[index_] != "" {
				if tmp, err := strconv.Atoi(_g[index_]); err == nil {
					data.additions = append(data.additions, TimeAddition{ n: tmp, unit: unit_})
				}
			}
		}
		if _g[13] != "" {
			if tmp, err := strconv.ParseFloat(_g[4], 64); err == nil {
				n_ := int(tmp * 1e9)
				data.additions = append(data.additions, TimeAddition{ n: n_, unit: "µsec"})
			}
		}

		pos += len(_g[0])
	} else if _g := p_pos.FindStringSubmatch(_s); len(_g) > 0 {
		// (first|last day of) first|next year|month|day
		day_flg_ := _g[2]
		pos_flg_ := _g[3]
		word_ := _g[5]
		month_name_ := _g[6]
		day_name_ := _g[7]

		a := TimeAddition{
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
		// set UTC
		loc, err_ := time.LoadLocation("UTC")
		if err_ != nil {
			return -1
		}
		data.setLocation(loc)
	} else {
		return -1
	}

	return pos
}

func move(data *TimeData, a *TimeAddition) {
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
func add(data *TimeData, a *TimeAddition) {
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
	data.normalize()
}

// Parse string to a time.Time variable
//func ParseTimeFormat(format string, base *time.Time) (*time.Time, *TimeData) {
func ParseTimeFormat(format string, base *time.Time) (*time.Time, int64) {
	s := strings.TrimSpace(strings.ToLower(format))
	if s == "" {
		return nil, -1
	}

	initRegexp()
	//defer freeRegexp()

	// data
	data := NewTimeData()

	if base == nil {
		t_ := time.Now()
		base = &t_
	}
	data.setFromTime(base)

	// convert datetime
	s = p_ignore.ReplaceAllString(s, ` `)
	s = p_isoformat.ReplaceAllString(s, `$1 $2 $4`)
	s = p_norm_sp.ReplaceAllString(s, ` `)

	// parse string
	s_len := len(s)
	pos := 0
	cnts := 0
	for cnts < 15 && pos < s_len {
		pos = scanFormat(data, s, pos)
		if pos < 0 {
			return nil, -1
		}
		cnts++
	}

	// additions
	a_len := len(data.additions)
	for i := 0; i < a_len; i++ {
		move(data, &data.additions[i])
	}
	//fmt.Print("[data] -> ",data,"\n")

	// result
	// res := time.Date(data.y, time.Month(data.m), data.d, data.h, data.i, data.s, data.us*1e3, time.Local)
	res := data.Time()

	// convert to unix time
	//t := res.Unix()
	//_, offset := res.Zone()
	//t += int64(offset)
	//t -= int64(data.z)
	t := res.Unix()

	// return &res, &data
	return res, t
}