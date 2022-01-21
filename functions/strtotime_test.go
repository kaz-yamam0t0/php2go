package functions

import (
	"github.com/stretchr/testify/assert"
	"testing"
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
	assert.Equal(t, int64(1570181040), Strtotime("-1year -13months -8week", 1640769840))                                      // 2019-10-08 09:24:00

	// php `strtotime` seems not to do well with "next" or "prev" format and timezone.
	//assert.Equal(t, int64(1640790000), Strtotime("next Thursday",1640769840)) // 2021-12-29 15:00:00
	//assert.Equal(t, int64(1640530800), Strtotime("last Monday",1640769840)) // 2021-12-26 15:00:00
	assert.Equal(t, int64(1640856240), Strtotime("next Thursday", 1640769840)) // 2021-12-30 09:24:00
	assert.Equal(t, int64(1640597040), Strtotime("last Monday", 1640769840))   // 2021-12-27 09:24:00
}
