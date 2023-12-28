package hash

/*
	217. 存在重复元素 2023/10/06 12:20 - 12:22
给你一个整数数组 nums 。如果任一值在数组中出现至少两次返回true；
如果数组中每个元素互不相同，返回 false 。


示例 1：
输入：nums = [1,2,3,1]
输出：true

示例 2：
输入：nums = [1,2,3,4]
输出：false

示例 3：
输入：nums = [1,1,1,3,3,4,3,2,4,2]
输出：true
*/

// 我的办法，哈希表 O(n),O(n)
func containsDuplicate(nums []int) bool {
	hashMap := make(map[int]int, len(nums))
	for _, v := range nums {
		hashMap[v]++
	}
	for _, v := range hashMap {
		if v >= 2 {
			return true
		}
	}
	return false
}

//更简单的办法，排序，检验下一个是否重复 O(NlogN),O(logN)
