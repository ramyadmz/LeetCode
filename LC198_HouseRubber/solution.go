package lc198

func rob(nums []int) int {
	currMax, prevMax := 0, nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		prevMax, currMax = max(nums[i]+currMax, prevMax), prevMax
	}
	return prevMax
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
