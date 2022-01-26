package functions

import (
	//"strconv"
	"time"
	php2go_date "github.com/kaz-yamam0t0/php2go/functions/date"
)

// Parses a time string according to a specified format
func DateCreateFromFormat(format string, s string) (*time.Time, error) {
	return php2go_date.ParseFormat(format, s)
}
