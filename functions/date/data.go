package date

// Relative Additions or Subtractions
type TimeAddition struct {
	n    int
	unit string

	pos     string
	month   int
	weekday int
	word    string
	day_flg string
}

// Time Data
type TimeData struct {
	y         int
	m         int
	d         int
	h         int
	i         int
	s         int
	us        int
	day       int // Weekday (actually doesn't affect the result)
	z         int
	additions []TimeAddition
}

