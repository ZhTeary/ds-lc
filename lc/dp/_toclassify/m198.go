package dp

/*198. 打家劫舍 medium 2023-12-15 21:55~22:14 19min */

//空间优化 - 滚动数组
func rov_v5(nums []int) int {

	first := nums[0]
	second := max(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		first, second = second, max(first+nums[i], second)
	}

	return second
}

//抄一遍题解后自己写的
func rob_v4(nums []int) int {
	n := len(nums)
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = nums[0]
	for i := 2; i < len(dp); i++ {
		dp[i] = max(dp[i-2]+nums[i-1], dp[i-1])
	}

	return dp[n]
}

//题解
/*!!!!	dp就比nums长1, 虽然别扭,
但减少 数! 组! 越! 界!  !!!!!!*/
func rob_v3(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}

	dp := make([]int, n+1)

	dp[0] = 0
	dp[1] = nums[0]

	for i := 2; i <= n; i++ { //这里i<len(dp)，就是i<=n
		dp[i] = max(dp[i-1], dp[i-2]+nums[i-1])
	}

	return dp[len(dp)-1]
}

//根据题解思路自己写的，比自己的思路自己写的还垃圾
func rob_v2(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	} else if len(nums) == 2 {

		return max(nums[0], nums[1])
	}

	money := 0
	dp := make([]int, len(nums))

	dp[0] = nums[0]
	dp[1] = nums[1]
	dp[2] = max(nums[0]+nums[2], nums[1])
	for i := 3; i < len(nums); i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i], dp[i-3]+nums[i])
	}

	for i := range dp {
		if dp[i] > money {
			money = dp[i]
		}
	}
	return money
}

//我的解法，丑陋
func rob_v1(nums []int) int {
	money := 0
	//特殊情况
	if len(nums) == 1 {
		return nums[0]
	} else if len(nums) == 2 {
		return max(nums[0], nums[1])
	} else if len(nums) == 3 {
		return max(nums[0]+nums[2], nums[1])
	}

	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = nums[1]
	dp[2] = nums[0] + nums[2]
	for i := 3; i < len(nums); i++ {
		dp[i] = max(nums[i]+dp[i-2], nums[i]+dp[i-3])
	}
	for i := range dp {
		if dp[i] > money {
			money = dp[i]
		}
	}
	return money
}
