package date

import "time"

const (
	AM = 1
	PM = 2
)
const (
	SET_YEAR              = 1
	SET_MONTH             = 2
	SET_DAY               = 4
	SET_HOUR              = 8
	SET_MINUTE            = 16
	SET_SECOND            = 32
	SET_MICROSECOND       = 64
	SET_TIMEZONE_OFFSET   = 128
	SET_TIMEZONE_LOCATION = 256

	SET_AP                = 512

	SKIP_ERRORS           = 1024
)


// Relative Additions or Subtractions
type TimeAddition struct {
	n    int     // number of difference
	unit string  // unit of difference

	pos     string // first | next
	month   int // number of month
	weekday int // number of weekday
	word    string // year | month|day
	day_flg string // empty | first | last (day of)
}

// Time Data
type TimeData struct {
	y         int // Year
	m         int // Month
	d         int // Date
	h         int // Hour
	i         int // Minute
	s         int // Second
	us        int // Microsecond
	ap        int // AM/PM flag (1=AM 2=PM)
	day       int // Weekday (actually doesn't affect the result)
	z         int // Timezone Offset
	loc       *time.Location // Timezone Location
	additions []TimeAddition // Relative differences

	flags     int // flags
}

// create new TimeData
func NewTimeData() *TimeData {
	d := TimeData{1970, 1, 1, 0, 0, 0, 0, 0, 0, 0, nil, []TimeAddition{}, 0}
	return &d
}

// resetIfUnset
func (data *TimeData) resetIfUnset() {
	if ! data.hasFlag(SET_YEAR) {
		data.y = 1970
	}
	if ! data.hasFlag(SET_MONTH) {
		data.m = 1
	}
	if ! data.hasFlag(SET_DAY) {
		data.d = 1
	}
	if ! data.hasFlag(SET_HOUR) {
		data.h = 0
	}
	if ! data.hasFlag(SET_MINUTE) {
		data.i = 0
	}
	if ! data.hasFlag(SET_SECOND) {
		data.s = 0
	}
	if ! data.hasFlag(SET_MICROSECOND) {
		data.us = 0
	}
	if ! data.hasFlag(SET_TIMEZONE_OFFSET) {
		data.z = 0
	}
	if ! data.hasFlag(SET_TIMEZONE_LOCATION) {
		data.loc = nil
	}
}
// reset
func (data *TimeData) reset() {
	data.y = 1970
	data.m = 1
	data.d = 1
	data.h = 0
	data.i = 0
	data.s = 0
	data.us = 0
	data.z = 0
	data.loc = nil

	data.flags &= (^SET_YEAR & ^SET_MONTH & ^SET_DAY & ^SET_HOUR & ^SET_MINUTE & ^SET_SECOND & ^SET_MICROSECOND & ^SET_TIMEZONE_OFFSET & ^SET_TIMEZONE_LOCATION)
}

// ============================================================
// setter
// ============================================================

func (data *TimeData) setYear(y int) {
	data.y = y
	data.flags |= SET_YEAR
}
func (data *TimeData) setMonth(m int) {
	data.m = m
	data.flags |= SET_MONTH
}
func (data *TimeData) setDay(d int) {
	data.d = d
	data.flags |= SET_DAY
}
func (data *TimeData) setHour(h int) {
	data.h = h
	data.flags |= SET_HOUR
}
func (data *TimeData) setMinute(i int) {
	data.i = i
	data.flags |= SET_MINUTE
}
func (data *TimeData) setSecond(s int) {
	data.s = s
	data.flags |= SET_SECOND
}
func (data *TimeData) setMicrosecond(s int) {
	data.us = s
	data.flags |= SET_MICROSECOND
}
func (data *TimeData) setTimezoneOffset(z int) {
	data.z = z
	data.flags |= SET_TIMEZONE_OFFSET
}
func (data *TimeData) setLocation(loc *time.Location) {
	data.loc = loc
	data.flags |= SET_TIMEZONE_LOCATION
}
func (data *TimeData) setFromTime(t *time.Time) {
	data.setYear(t.Year())
	data.setMonth(int(t.Month()))
	data.setDay(t.Day())
	data.setHour(t.Hour())
	data.setMinute(t.Minute())
	data.setSecond(t.Second())
	data.setMicrosecond(int(t.Nanosecond() / 1e3))
	data.setLocation(t.Location())
	
	//_, offset_ := base.Zone()
	data.setTimezoneOffset(0)
}

// ============================================================
// flags
// ============================================================
func (data *TimeData) hasFlag(f int) bool {
	return (data.flags & f) == f
}


// ============================================================
// normalize
// ============================================================

// normalize Year, Month, Date, Hour, Minute, Second, Millisecond
func (data *TimeData) normalize() {
	data.s, data.us = norm(data.s, data.us, 1000000)
	data.i, data.s = norm(data.i, data.s, 60)
	data.h, data.i = norm(data.h, data.i, 60)
	data.d, data.h = norm(data.d, data.h, 24)

	// day, month, year
	data.normalizeYmd()
}

// normalize Year, Month, Date
func (data *TimeData) normalizeYmd() {
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

// ============================================================
// time
// ============================================================

func (data *TimeData) Time() *time.Time {
	loc := time.Local
	if data.loc != nil {
		loc = data.loc
	}
	res := time.Date(data.y, time.Month(data.m), data.d, data.h, data.i, data.s, data.us*1e3, loc)

	if data.z != 0 {
		tmp := res.Unix()
		tmp += int64(data.z) 
		res = time.Unix(tmp,0).In(loc)
	}

	return &res
}
