package hash

/*
	1207. 独一无二的出现次数  简单   2023/10/03 11:24 - 11:30

给你一个整数数组 arr，请你帮忙统计数组中每个数的出现次数。

如果每个数的出现次数都是独一无二的，就返回 true；否则返回 false。



示例 1：

输入：arr = [1,2,2,1,1,3]
输出：true
解释：在该数组中，1 出现了 3 次，2 出现了 2 次，3 只出现了 1 次。没有两个数的出现次数相同。
示例 2：

输入：arr = [1,2]
输出：false
示例 3：

输入：arr = [-3,0,1,-3,1,1,1,-3,10,0]
输出：true
*/

func UniqueOccurrences(arr []int) bool {
	hashMap := make(map[int]int)
	for _, v := range arr {
		hashMap[v]++
	}
	var ans []int
	for _, v := range hashMap {
		ans = append(ans, v)
	}

	for i := 0; i < len(ans); i++ {
		for j := i + 1; j < len(ans); j++ {
			if ans[i] == ans[j] {
				return false
			}
		}
	}
	return true

}
