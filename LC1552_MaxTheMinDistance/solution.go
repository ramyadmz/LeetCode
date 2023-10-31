package lc1552maxthemindistance

import "sort"

func maxDistance(position []int, m int) int {
	var ans int
	sort.Ints(position)
	n := len(position)
	low, high := 1, position[n-1]-position[0]

	for low <= high {
		mid := low + (high-low)/2
		if func(dist int) bool {
			prev, count := position[0], 1
			for _, curr := range position[1:] {
				if curr-prev >= dist {
					prev = curr
					count++
				}
			}
			return count >= m
		}(mid) {
			ans = mid
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return ans
}
