package lc1360

import "strconv"

func DaysBetweenDates(date1 string, date2 string) int {
	y1, _ := strconv.Atoi(date1[0:4])
	m1, _ := strconv.Atoi(date1[5:7])
	d1, _ := strconv.Atoi(date1[8:10])

	y2, _ := strconv.Atoi(date2[0:4])
	m2, _ := strconv.Atoi(date2[5:7])
	d2, _ := strconv.Atoi(date2[8:10])

	return abs(countDays(y1, m1, d1) - countDays(y2, m2, d2))
}

func countDays(y, m, d int) (count int) {
	if m < 3 {
		m += 12
		y -= 1
	}
	daysInLeap := (y / 4) - (y / 100) + (y / 400)
	dayInYears := y * 365
	dayInMonths := (153*m + 8) / 5
	count = dayInYears + daysInLeap + dayInMonths + d
	return
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
