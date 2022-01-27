package functions

import (
	//"strconv"
	"github.com/kaz-yamam0t0/go-timeparser/timeparser"
	"time"
)

// Parses a time string according to a specified format
func DateCreateFromFormat(format string, s string) (*time.Time, error) {
	return timeparser.ParseFormat(format, s)
}
