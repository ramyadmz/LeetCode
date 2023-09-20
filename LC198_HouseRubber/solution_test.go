package lc198

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		houses   []int
		expected int
	}{
		{[]int{2, 1, 1, 2}, 4},
		{[]int{2, 7, 9, 3, 1}, 12},
	}
	for i := range tests {
		actual := rob(tests[i].houses)
		if actual != tests[i].expected {
			t.Errorf("test %d was unsuccessful : expected %v got %v", i, tests[i].expected, actual)
		}
	}

}
