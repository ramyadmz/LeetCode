package lc946

func validateStackSequences(pushed []int, popped []int) bool {
	j, top := 0, -1
	for i := range pushed {
		top++
		pushed[top] = pushed[i]
		for top >= 0 && pushed[top] == popped[j] {
			top--
			j++
		}
	}
	return top == -1
}
