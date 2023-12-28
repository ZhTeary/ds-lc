package hash

/*
	2215. 找出两组数的不同 简单  2023-10-25 10:34-10:46  X  case2没过

给你两个下标从 0 开始的整数数组 nums1 和 nums2 ，请你返回一个长度为 2 的列表 answer ，其中：

answer[0] 是 nums1 中所有 不 存在于 nums2 中的 不同 整数组成的列表。
answer[1] 是 nums2 中所有 不 存在于 nums1 中的 不同 整数组成的列表。
注意：列表中的整数可以按 任意 顺序返回。


示例 1：

输入：nums1 = [1,2,3], nums2 = [2,4,6]
输出：[[1,3],[4,6]]
解释：
对于 nums1 ，nums1[1] = 2 出现在 nums2 中下标 0 处，然而 nums1[0] = 1 和 nums1[2] = 3 没有出现在 nums2 中。因此，answer[0] = [1,3]。
对于 nums2 ，nums2[0] = 2 出现在 nums1 中下标 1 处，然而 nums2[1] = 4 和 nums2[2] = 6 没有出现在 nums2 中。因此，answer[1] = [4,6]。
示例 2：

输入：nums1 = [1,2,3,3], nums2 = [1,1,2,2]
输出：[[3],[]]
解释：
对于 nums1 ，nums1[2] 和 nums1[3] 没有出现在 nums2 中。由于 nums1[2] == nums1[3] ，二者的值只需要在 answer[0] 中出现一次，故 answer[0] = [3]。
nums2 中的每个整数都在 nums1 中出现，因此，answer[1] = [] 。
*/

/*
解题思路：
1.将两个数组转变为map，这样就可以避免数组中的重复
2.在两个map中互相循环，判断元素是否存在与另一个数组中
3.时间复杂度为O(m+n)
*/
func findDifference(nums1 []int, nums2 []int) [][]int {
	//map 用make创建
	hashMap1 := make(map[int]bool)
	hashMap2 := make(map[int]bool)
	res := make([][]int, 2)

	for _, num := range nums1 {
		hashMap1[num] = true
	}
	for _, num := range nums2 {
		hashMap2[num] = true
	}

	//对于map的循环，循环里的元素是key
	for num := range hashMap1 {
		if _, has := hashMap2[num]; !has {
			res[0] = append(res[0], num)
		}
	}

	for num := range hashMap2 {
		if _, has := hashMap1[num]; !has {
			res[1] = append(res[1], num)
		}
	}

	return res
}

// 错误1 不应该用数组，golang虽然没有set,但是可以用map[]bool
func findDifferencewrong(nums1 []int, nums2 []int) [][]int {
	var ans1 []int
	var ans2 []int
	flag := true
	for _, v := range nums1 {
		for _, u := range nums2 {
			if v == u {
				flag = false
			}
		}
		if flag {
			ans1 = append(ans1, v)
		}
		flag = true
	}

	for _, v := range nums2 {
		for _, u := range nums1 {
			if v == u {
				flag = false
			}
		}
		if flag {
			ans2 = append(ans2, v)
		}
		flag = true
	}

	var ans [][]int
	ans = append(ans, ans1, ans2)
	return ans
}
