package functions

import (
	"time"
	php2go_date "github.com/kaz-yamam0t0/php2go/functions/date"
)


func Str2Time(format string, base *time.Time) *time.Time {
	res, _ := php2go_date.ParseTimeFormat(format, base)
	return res
}

// Parse about any English textual datetime description into a Unix timestamp
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
	res, t := php2go_date.ParseTimeFormat(format, &base)
	if res == nil {
		return -1
	}

	return t
}
