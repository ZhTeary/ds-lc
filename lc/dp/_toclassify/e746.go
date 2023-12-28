package dp

func minCostClimbingStairs(cost []int) int {
	n := len(cost)
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		dp[i] = mini(dp[n-1]+cost[n-1], dp[n-2]+cost[n-2])
	}

	return dp[n]

}

func mini(a, b int) int {
	if a > b {
		return b
	}
	return a
}
