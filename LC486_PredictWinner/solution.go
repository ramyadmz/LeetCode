package lc486

func predictTheWinner(nums []int) bool {
	n := len(nums)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	var helper func(start, end int) int
	helper = func(start, end int) int {
		if dp[start][end] != 0 {
			return dp[start][end]
		}
		if start == end {
			dp[start][end] = nums[start]
			return nums[start]
		}
		dp[start][end] = max(nums[start]-helper(start+1, end), nums[end]-helper(start, end-1))
		return dp[start][end]
	}
	return helper(0, n-1) >= 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
