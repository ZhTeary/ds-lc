package hash

import "sort"

/*
 169. 多数元素 简单

    2023-10-06 第一次 3分 时间O(n) 空间O(n)

给定一个大小为 n 的数组 nums ，返回其中的多数元素。多数元素是指在数组中出现次数 大于 ⌊ n/2 ⌋ 的元素。

你可以假设数组是非空的，并且给定的数组总是存在多数元素。

示例 1：
输入：nums = [3,2,3]
输出：3

示例 2：
输入：nums = [2,2,1,1,1,2,2]
输出：2

进阶：尝试设计时间复杂度为 O(n)、空间复杂度为 O(1) 的算法解决此问题。
*/
func majorityElement(nums []int) int {

	sort.Ints(nums)
	count := 0
	maxi := nums[0]
	for i := range nums {
		if nums[i] == maxi {
			count++
			if count > len(nums)/2 {
				return nums[i]
			}
		} else {
			count = 1
			maxi = nums[i]
		}
	}
	return 0
}

//func MajorityElement(nums []int) int {
//
//	hashMap := make(map[int]int, len(nums))
//	for _, v := range nums {
//		hashMap[v]++
//	}
//	n := len(nums) / 2
//	for i, v := range hashMap {
//		if v > n {
//			return i
//		}
//	}
//	return 0
//}
