package lc1552maxthemindistance

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		list     []int
		balls    int
		expected int
	}{
		{[]int{1, 2, 3, 4, 7}, 3, 3},
	}
	for _, test := range tests {
		if actual := maxDistance(test.list, test.balls); actual != test.expected {
			t.Errorf("expected: %d got %d", test.expected, actual)
		}
	}
}
