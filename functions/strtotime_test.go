package functions

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestStrtotime(t *testing.T) {

	// General Format
	//assert.Equal(t, 1640769840, Strtotime("now"))
	assert.Equal(t, int64(968511600), Strtotime("10 September 2000"))
	assert.Equal(t, int64(968511600), Strtotime("2000/09/10"))
	assert.Equal(t, int64(968511600), Strtotime("2000-09-10"))
	assert.Equal(t, int64(968556896), Strtotime("2000-09-10 12:34:56"))
	assert.Equal(t, int64(1640769840), Strtotime("Wed, 29 Dec 2021 18:24:00 +0900"))
	assert.Equal(t, int64(1640773440), Strtotime("Wed, 29 Dec 2021 18:24:00 +0800"))
	assert.Equal(t, int64(1640769840), Strtotime("2021-12-29T18:24:00+09:00"))
	assert.Equal(t, int64(1640769840), Strtotime("Wednesday 29th December 2021 06:24:00 PM"))

	// American or European format (Ambiguious)
	assert.Equal(t, int64(968511600), Strtotime("9/10/2000"))
	assert.Equal(t, int64(968511600), Strtotime("10.9.2000"))
	assert.Equal(t, int64(968511600), Strtotime("10-9-2000"))
	//assert.Equal(t, 968511600, Strtotime("2000.9.10")) // returns false.

	// Relative Format
	assert.Equal(t, int64(1640769840), Strtotime("+0 day", 1640769840))                                                       // 2021-12-29 09:24:00
	assert.Equal(t, int64(1640856240), Strtotime("+1 day", 1640769840))                                                       // 2021-12-30 09:24:00
	assert.Equal(t, int64(1641461040), Strtotime("+1 day 1week", 1640769840))                                                 // 2022-01-06 09:24:00
	assert.Equal(t, int64(1523629087), Strtotime("+1 day 2 week 3 months -4 years 5 hours -6 minutes 7 seconds", 1640769840)) // 2018-04-13 14:18:07
	assert.Equal(t, int64(1570181040), Strtotime("-1year -13months -8weeks", 1640769840))                                     // 2019-10-08 09:24:00
	assert.Equal(t, int64(1570181040), Strtotime("1 year ago 13months ago 8 weeks ago", 1640769840))                          // 2019-10-08 09:24:00
	
	assert.Equal(t, int64(1640856240), Strtotime("P1D", 1640769840))                                                       // 2021-12-30 09:24:00
	
	// php `strtotime` seems not to do well with "next" or "prev" format and timezone.
	//assert.Equal(t, int64(1640790000), Strtotime("next Thursday",1640769840)) // 2021-12-29 15:00:00
	//assert.Equal(t, int64(1640530800), Strtotime("last Monday",1640769840)) // 2021-12-26 15:00:00
	assert.Equal(t, int64(1640856240), Strtotime("next Thursday", 1640769840)) // 2021-12-30 09:24:00
	assert.Equal(t, int64(1640597040), Strtotime("last Monday", 1640769840))   // 2021-12-27 09:24:00
}
func ExampleStrtotime() {
	// Strtotime(format string) returns int64
	// or -1 when an error has occurred.

	// general format
	_ = Strtotime("10 September 2000")
	_ = Strtotime("2000/09/10")
	_ = Strtotime("2000-09-10")
	_ = Strtotime("2000-09-10 12:34:56")

	// American or European format (Ambiguious)
	_ = Strtotime("9/10/2000")
	_ = Strtotime("10.9.2000")
	_ = Strtotime("10-9-2000")
	// _ = Strtotime("2000.9.10") error

	// various formats
	_ = Strtotime("Wed, 29 Dec 2021 18:24:00 +0900")
	_ = Strtotime("2021-12-29T18:24:00+09:00")
	_ = Strtotime("Wednesday 29th December 2021 06:24:00 PM")

	// relative format with int64 unixtime (seconds) or time.Time variable.
	_ = Strtotime("+0 day", time.Now()) // import "time"
	_ = Strtotime("+1 day 2 week 3 months -4 years 5 hours -6 minutes 7 seconds", time.Now())
	_ = Strtotime("-1year -13months -8week", time.Now())
	_ = Strtotime("1 year ago", time.Now())
	_ = Strtotime("P1Y2M3DT4H5M6.123456S", time.Now()) // ISO 8601 Interval Format

	_ = Strtotime("next Thursday", time.Now())
	_ = Strtotime("last Monday", time.Now())
}

