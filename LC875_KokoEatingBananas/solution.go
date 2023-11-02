package lc875kokoeatingbananas

func minEatingSpeed(piles []int, h int) int {
	possible := func(speed int) bool {
		var count int
		for _, pile := range piles {
			count += (pile + speed - 1) / speed
		}
		return count <= h
	}

	low, high := 1, 1000000000
	// instead of finding max or sorting the list we use max constrant as with binary search it take couple of steps to reach the small number
	for low < high {
		mid := low + (high-low)/2
		if possible(mid) {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return low
}
