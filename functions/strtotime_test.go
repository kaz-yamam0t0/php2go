package functions

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestStrtotime(t *testing.T) {
	utc, err := time.LoadLocation("UTC")
	assert.Nil(t, err)

	// General Format
	//assert.Equal(t, 1640769840, Strtotime("now"))
	assert.Equal(t, time.Date(2000, time.September, 10, 0, 0, 0, 0, time.Local).Unix(), Strtotime("10 September 2000"))
	assert.Equal(t, time.Date(2000, time.September, 10, 0, 0, 0, 0, time.Local).Unix(), Strtotime("2000/09/10"))
	assert.Equal(t, time.Date(2000, time.September, 10, 0, 0, 0, 0, time.Local).Unix(), Strtotime("2000-09-10"))
	assert.Equal(t, time.Date(2000, time.September, 10, 12, 34, 56, 0, time.Local).Unix(), Strtotime("2000-09-10 12:34:56"))
	assert.Equal(t, time.Date(2021, time.December, 29, 18, 24, 0, 0, utc).Unix(), Strtotime("Wed, 29 Dec 2021 18:24:00 +0000"))
	assert.Equal(t, time.Date(2021, time.December, 30, 3, 24, 0, 0, utc).Unix(), Strtotime("Wed, 29 Dec 2021 18:24:00 +0900"))
	assert.Equal(t, time.Date(2021, time.December, 30, 2, 24, 0, 0, utc).Unix(), Strtotime("Wed, 29 Dec 2021 18:24:00 +08:00"))
	assert.Equal(t, time.Date(2021, time.December, 30, 3, 24, 0, 0, utc).Unix(), Strtotime("2021-12-29T18:24:00+09:00"))
	assert.Equal(t, time.Date(2021, time.December, 29, 18, 24, 0, 0, time.Local).Unix(), Strtotime("Wednesday 29th December 2021 06:24:00 PM"))

	// American or European format (Ambiguious)
	assert.Equal(t, time.Date(2000, time.September, 10, 0, 0, 0, 0, time.Local).Unix(), Strtotime("9/10/2000"))
	assert.Equal(t, time.Date(2000, time.September, 10, 0, 0, 0, 0, time.Local).Unix(), Strtotime("10.9.2000"))
	assert.Equal(t, time.Date(2000, time.September, 10, 0, 0, 0, 0, time.Local).Unix(), Strtotime("10-9-2000"))
	//assert.Equal(t, 968511600, Strtotime("2000.9.10")) // returns false.

	// Relative Format
	tm := time.Date(2000, time.September, 10, 0, 0, 0, 0, time.Local)
	assert.Equal(t, time.Date(2000, time.September, 10, 0, 0, 0, 0, time.Local).Unix(), Strtotime("+0 day", tm.Unix()))
	assert.Equal(t, time.Date(2000, time.September, 11, 0, 0, 0, 0, time.Local).Unix(), Strtotime("+1 day", tm.Unix()))
	assert.Equal(t, time.Date(2000, time.September, 18, 0, 0, 0, 0, time.Local).Unix(), Strtotime("+1 day 1week", tm.Unix()))                                                 // 2022-01-06 09:24:00
	assert.Equal(t, time.Date(1996, time.December, 25, 4, 54, 7, 0, time.Local).Unix(), Strtotime("+1 day 2 week 3 months -4 years 5 hours -6 minutes 7 seconds", tm.Unix())) // 2018-04-13 14:18:07
	assert.Equal(t, time.Date(1998, time.June, 15, 0, 0, 0, 0, time.Local).Unix(), Strtotime("-1year -13months -8weeks", tm.Unix()))                                          // 2019-10-08 09:24:00
	assert.Equal(t, time.Date(1998, time.June, 15, 0, 0, 0, 0, time.Local).Unix(), Strtotime("1 year ago 13months ago 8 weeks ago", tm.Unix()))                               // 2019-10-08 09:24:00

	assert.Equal(t, time.Date(2000, time.September, 11, 0, 0, 0, 0, time.Local).Unix(), Strtotime("P1D", tm.Unix()))               // 2021-12-30 09:24:00
	assert.Equal(t, time.Date(2001, time.November, 13, 0, 0, 0, 0, time.Local).Unix(), Strtotime("P1Y2M3D", tm.Unix()))            // 2021-12-30 09:24:00
	assert.Equal(t, time.Date(2001, time.November, 13, 4, 5, 6, 0, time.Local).Unix(), Strtotime("P1Y2M3DT4H5M6S", tm.Unix()))     // 2021-12-30 09:24:00
	assert.Equal(t, time.Date(2001, time.November, 13, 4, 5, 6, 7e8, time.Local).Unix(), Strtotime("P1Y2M3DT4H5M6.7S", tm.Unix())) // 2021-12-30 09:24:00
	//

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
	fmt.Println(Strtotime("10 September 2000"))
	fmt.Println(Strtotime("2000/09/10"))
	fmt.Println(Strtotime("2000-09-10"))
	fmt.Println(Strtotime("2000-09-10 12:34:56"))

	// American or European format (Ambiguous)
	fmt.Println(Strtotime("9/10/2000"))
	fmt.Println(Strtotime("10.9.2000"))
	fmt.Println(Strtotime("10-9-2000"))
	// _ = Strtotime("2000.9.10") error

	// various formats
	fmt.Println(Strtotime("Wed, 29 Dec 2021 18:24:00 +0900"))
	fmt.Println(Strtotime("2021-12-29T18:24:00+09:00"))
	fmt.Println(Strtotime("Wednesday 29th December 2021 06:24:00 PM"))

	// relative format with int64 unixtime (seconds) or time.Time variable.
	fmt.Println(Strtotime("+0 day", time.Now())) // import "time"
	fmt.Println(Strtotime("+1 day 2 week 3 months -4 years 5 hours -6 minutes 7 seconds", time.Now()))
	fmt.Println(Strtotime("-1year -13months -8week", time.Now()))
	fmt.Println(Strtotime("1 year ago", time.Now()))
	fmt.Println(Strtotime("P1Y2M3DT4H5M6.123456S", time.Now())) // ISO 8601 Interval Format

	fmt.Println(Strtotime("next Thursday", time.Now()))
	fmt.Println(Strtotime("last Monday", time.Now()))
}
