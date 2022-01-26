package functions

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// import "time"

func TestDateCreateFromFormat(t *testing.T) {
	cases := []string{
		" l jS \\of F Y  h:i:s A",
		"Wednesday 29th of December  2021 06:24:12 PM ",

		"  Y-m-d    H:i:s ",
		" 2021-12-29  18:24:12  ",

		"Y#n#j H:i:s",
		"2021,12,29 18:24:12",

		"U H?i#s",
		"1640769840 18_24:12",
	}
	l := len(cases)
	for i := 0; i < l; i += 2 {
		tm, err := DateCreateFromFormat(cases[i], cases[i+1])

		assert.Nil(t, err)
		assert.Equal(t, 2021, tm.Year())
		assert.Equal(t, 12, int(tm.Month()))
		assert.Equal(t, 29, tm.Day())
		assert.Equal(t, 18, tm.Hour())
		assert.Equal(t, 24, tm.Minute())
		assert.Equal(t, 12, tm.Second())
	}

	cases = []string{
		"Y?n*j|",
		"2021a12___29 18:24:12",

		"Y#n#j|+j",
		"2021/12/29 ___",
	}
	l = len(cases)
	for i := 0; i < l; i += 2 {
		tm, err := DateCreateFromFormat(cases[i], cases[i+1])

		assert.Nil(t, err)
		assert.Equal(t, 2021, tm.Year())
		assert.Equal(t, 12, int(tm.Month()))
		assert.Equal(t, 29, tm.Day())
		assert.Equal(t, 0, tm.Hour())
		assert.Equal(t, 0, tm.Minute())
		assert.Equal(t, 0, tm.Second())
	}

	// Timezone
	cases = []string{
		"UTC",
		"GMT",
		"MST",
		//"JST", // "Asia/Tokyo",

		"Africa/Johannesburg",
		"Africa/Lagos",
		"Africa/Windhoek",
		"America/Adak",
		"America/Bogota",
		"America/Caracas",
		"America/Chicago",
		"America/Denver",
		"America/Godthab",
		"America/Guatemala",
		"America/Halifax",
		"America/Los_Angeles",
		"America/Montevideo",
		"America/New_York",
		"America/Noronha",
		"America/Phoenix",
		"America/Santiago",
		"America/Santo_Domingo",
		"America/St_Johns",
		"Asia/Baghdad",
		"Asia/Baku",
		"Asia/Beirut",
		"Asia/Dhaka",
		"Asia/Dubai",
		"Asia/Irkutsk",
		"Asia/Jakarta",
		"Asia/Kabul",
		"Asia/Kamchatka",
		"Asia/Karachi",
		"Asia/Kathmandu",
		"Asia/Kolkata",
		"Asia/Krasnoyarsk",
		"Asia/Omsk",
		"Asia/Rangoon",
		"Asia/Shanghai",
		"Asia/Tehran",
		"Asia/Tokyo",
		"Asia/Vladivostok",
		"Asia/Yakutsk",
		"Asia/Yekaterinburg",
		"Atlantic/Azores",
		"Atlantic/Cape_Verde",
		"Australia/Adelaide",
		"Australia/Brisbane",
		"Australia/Darwin",
		"Australia/Eucla",
		"Australia/Lord_Howe",
		"Australia/Sydney",
		"Etc/GMT+12",
		"Europe/Berlin",
		"Europe/London",
		"Europe/Moscow",
		"Pacific/Apia",
		"Pacific/Auckland",
		"Pacific/Chatham",
		"Pacific/Easter",
		"Pacific/Gambier",
		"Pacific/Honolulu",
		"Pacific/Kiritimati",
		"Pacific/Majuro",
		"Pacific/Marquesas",
		"Pacific/Norfolk",
		"Pacific/Noumea",
		"Pacific/Pago_Pago",
		"Pacific/Pitcairn",
		"Pacific/Tongatapu",
	}
	l = len(cases)
	for i := 0; i < l; i++ {
		tm, err := DateCreateFromFormat("  Y-m-d    H:i:s T", " 2021-12-29  18:24:12 "+cases[i])

		assert.Nil(t, err)
		assert.Equal(t, 2021, tm.Year())
		assert.Equal(t, 12, int(tm.Month()))
		assert.Equal(t, 29, tm.Day())
		assert.Equal(t, 18, tm.Hour())
		assert.Equal(t, 24, tm.Minute())
		assert.Equal(t, 12, tm.Second())

		assert.Equal(t, cases[i], tm.Location().String())
	}
	// offset
	// Timezone
	cases_i := []int{
		0,
		-1,
		-11,
		+1,
		+11,
	}
	l = len(cases_i)
	for i := 0; i < l; i++ {
		n_ := cases_i[i]
		tm, err := DateCreateFromFormat("  Y-m-d    H:i:s T", fmt.Sprintf(" 2021-12-29  12:24:12 %+02d:00", n_))

		assert.Nil(t, err)
		assert.Equal(t, 2021, tm.Year())
		assert.Equal(t, 12, int(tm.Month()))
		assert.Equal(t, 29, tm.Day())
		assert.Equal(t, 12+n_, tm.Hour())
		assert.Equal(t, 24, tm.Minute())
		assert.Equal(t, 12, tm.Second())

		assert.Equal(t, "UTC", tm.Location().String())
	}
}

func ExampleDateCreateFromFormat() {
	// `DateCreateFromFormat` returns a time.Time variable
	// 2021-12-29 18:24:12 +0900 JST
	tm, err := DateCreateFromFormat(" l jS \\of F Y  h:i:s A", "Wednesday 29th of December  2021 06:24:12 PM ")
	if err != nil {
		panic(err)
	}
	fmt.Println(tm)

	// Timezone
	// 2021-12-29 06:24:12 +0900 JST
	tm, err = DateCreateFromFormat("Y-m-d H:i:s T", "2021-12-29 06:24:12 Asia/Tokyo")
	if err != nil {
		panic(err)
	}
	fmt.Println(tm)

	// Timezone Offset means add or subtract offset with UTC Timezone
	// 2021-12-29 15:24:12 +0000 UTC
	tm, err = DateCreateFromFormat("Y-m-d H:i:s T", "2021-12-29 06:24:12 +09:00")
	if err != nil {
		panic(err)
	}
	fmt.Println(tm)
}
