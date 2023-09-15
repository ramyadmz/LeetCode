package lc486

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		nums   []int
		expect bool
	}{
		{[]int{1, 5, 233, 7}, true},
		{[]int{1, 5, 2}, false},
	}

	for i, test := range tests {
		actual := predictTheWinner(test.nums)
		if test.expect != actual {
			t.Errorf("test %d was unsuccessful : expected %v got %v", i, test.expect, actual)

		}
	}

}
