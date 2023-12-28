package hash

/*
 219. 存在重复元素II 简单 2023/10/06 12:25-12:40

给你一个整数数组 nums 和一个整数 k ,
判断数组中是否存在两个 不同的索引 i 和 j,
满足 nums[i] == nums[j] 且 abs(i - j) <= k 。
如果存在，返回 true ；否则，返回 false 。

示例 1：
输入：nums = [1,2,3,1], k = 3
输出：true

示例 2：
输入：nums = [1,0,1,1], k = 1
输出：true

示例 3：
输入：nums = [1,2,3,1,2,3], k = 2
输出：false
*/

//先判断哪些元素有重复，再判断哪些距离小

func containsNearbyDuplicate(nums []int, k int) bool {
	for i := range nums {
		for j := i + 1; j <= i+k && j < len(nums); j++ {
			if nums[j] == nums[i] {
				return true
			}
		}
	}
	return false
}
