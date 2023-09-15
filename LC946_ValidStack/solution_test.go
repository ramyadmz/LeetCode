package lc946

import "testing"

func Test(t *testing.T) {
	tests := []struct {
		pushed   []int
		popped   []int
		expected bool
	}{
		{[]int{1, 2, 3, 4, 5}, []int{4, 5, 3, 2, 1}, true},
		{[]int{1, 2, 3, 4, 5}, []int{4, 3, 5, 1, 2}, false},
	}

	for i := range tests {
		result := validateStackSequences(tests[i].pushed, tests[i].popped)
		if result != tests[i].expected {
			t.Errorf(" test case %d faild: expected %v got %v", i, tests[i].expected, result)
		}
	}

}
