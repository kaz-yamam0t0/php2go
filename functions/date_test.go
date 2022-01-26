package functions

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)
import "time"

func TestDate(t *testing.T) {
	d := time.Unix(int64(1640769840), 0)

	tz, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}
	d = d.In(tz)

	// The second argument can be both int or a `time.Time` object.
	assert.Equal(t, "Wed, 29 Dec 2021 18:24:00 +0900", Date("r", 1640769840))
	assert.Equal(t, "Wed, 29 Dec 2021 18:24:00 +0900", Date("r", d))

	// Each format
	assert.Equal(t, "Wednesday 29th of December 2021 06:24:00 PM", Date("l jS \\of F Y h:i:s A", d))
	assert.Equal(t, "2021-12-29 18:24:00", Date("Y-m-d H:i:s", d))
	assert.Equal(t, "21 12 29", Date("y n j", d))
	assert.Equal(t, "52", Date("W", d))
	assert.Equal(t, "Wed Wednesday 3 3", Date("D l w N", d))       // Day
	assert.Equal(t, "December 12 Dec 12 31", Date("F m M n t", d)) // Month
	assert.Equal(t, "2021 21 0 2021", Date("Y y L o", d))          // Year
	assert.Equal(t, "pm PM 433", Date("a A B", d))                 // Time
	assert.Equal(t, "6 18 06 18", Date("g G h H", d))              // Hour
	assert.Equal(t, "24 00 000000 000", Date("i s u v", d))        // Minutes / Seconds
	assert.Equal(t, "32400 +09:00 +0900", Date("Z P O", d))        // Timezone
	assert.Equal(t, "2021-12-29T18:24:00+09:00", Date("c", d))
	assert.Equal(t, "JST", Date("e", d))
	assert.Equal(t, "JST", Date("T", d))
	assert.Equal(t, "32400", Date("Z", d))
	assert.Equal(t, "+09:00", Date("P", d))
	assert.Equal(t, "+0900", Date("O", d))
	assert.Equal(t, "1640769840", Date("U", d)) // Timestamp
	assert.Equal(t, "Wed, 29 Dec 2021 18:24:00 +0900", Date("r", d))
}

func ExampleDate() {
	// Date(format string, baseTime int*|time.Time) returns string
	// format -> https://www.php.net/manual/en/datetime.format.php

	fmt.Println(Date("r", time.Now()))                     // Wed, 29 Dec 2021 18:24:00 +0900
	fmt.Println(Date("l jS \\of F Y h:i:s A", time.Now())) // Wednesday 29th of December 2021 06:24:00 PM
	fmt.Println(Date("Y-m-d H:i:s", time.Now()))           // 2021-12-29 18:24:00
}
