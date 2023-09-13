package lc1360

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		date1  string
		date2  string
		result int
	}{
		{"2019-06-29", "2019-06-30", 1},
		{"2020-01-01", "2020-01-02", 1},
		{"2019-12-31", "2020-01-01", 1},
		{"2020-02-28", "2020-02-29", 1}, // Leap year
		{"2019-12-31", "2020-01-15", 15},
	}
	for _, test := range tests {
		if actual := DaysBetweenDates(test.date1, test.date2); actual != test.result {
			t.Errorf("expected %d got %d", test.result, actual)
		}
	}

}
