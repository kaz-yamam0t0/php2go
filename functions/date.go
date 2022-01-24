package functions

import (
	//"strconv"
	"time"
	php2go_date "github.com/kaz-yamam0t0/php2go/functions/date"
)

// Format a local time/date
// 
// Actually the sypnosis of Date() is like this:
// 
//  Date(s string, baseTime int*|time.Time)
func Date(s string, args ...interface{}) string {
	var dt time.Time

	// support both unixtime and time.Time object
	if len(args) > 0 {
		switch tmp := args[0].(type) {
		case int:
			dt = time.Unix(int64(tmp), 0)
		case int8:
			dt = time.Unix(int64(tmp), 0)
		case int16:
			dt = time.Unix(int64(tmp), 0)
		case int32:
			dt = time.Unix(int64(tmp), 0)
		case int64:
			dt = time.Unix(tmp, 0)
		case time.Time:
			dt = tmp
		default:
			panic("invalid arguments")
		}
	} else {
		dt = time.Now()
	}

	return php2go_date.DateFormat(s, dt)
}
