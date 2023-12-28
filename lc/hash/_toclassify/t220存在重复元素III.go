package hash

/*
 220. 存在重复元素 III 困难 2023/10/06 12:43

给你一个整数数组 nums 和两个整数 indexDiff 和 valueDiff 。

找出满足下述条件的下标对 (i, j)：
- i != j,
- abs(i - j) <= indexDiff
- abs(nums[i] - nums[j]) <= valueDiff
如果存在，返回 true ；否则，返回 false 。

示例 1：
输入：nums = [1,2,3,1], indexDiff = 3, valueDiff = 0
输出：true
解释：可以找出 (i, j) = (0, 3) 。
满足下述 3 个条件：
i != j --> 0 != 3
abs(i - j) <= indexDiff --> abs(0 - 3) <= 3
abs(nums[i] - nums[j]) <= valueDiff --> abs(1 - 1) <= 0

示例 2：
输入：nums = [1,5,9,1,5,9], indexDiff = 2, valueDiff = 3
输出：false
解释：尝试所有可能的下标对 (i, j) ，均无法满足这 3 个条件，因此返回 false 。
*/

// wrong1 超时
func containsNearbyAlmostDuplicate_times(nums []int, indexDiff int, valueDiff int) bool {

	for i := range nums {
		for j := i + 1; j <= i+indexDiff && j < len(nums); j++ {
			if abs(nums[i], nums[j]) <= valueDiff {
				return true
			}
		}
	}
	return false
}

func abs(i, j int) int {
	if i >= j {
		return i - j
	} else {
		return 0 - i + j
	}
}

//自我分析： 为什么超时？ 双重循环  怎么解决？空间换时间  空间怎么换时间？拆循环  用什么代替？
//思路： 把两个循环拆开，变为分别循环两次。 第一个循环将想要的信息都存储起来，第二个循环遍历新ds

// 先遍历一遍数组，value用出现间隔代替？ 存一个前一个满足的index，

// 两种解法，第一种是滑动窗口+有序集合 有序集合得自己写数据结构
// 第二种是桶排序
// 第二种：
// getID 是获取桶的编号?
func getID(x, w int) int {
	if x >= 0 {
		return x / w
	}
	return (x+1)/w - 1
}

func containsNearbyAlmostDuplicate(nums []int, k int, t int) bool {
	mp := make(map[int]int, len(nums))

	for i, x := range nums {
		id := getID(x, t+1)
		if _, has := mp[id]; has {
			return true
		}
		if y, has := mp[id-1]; has && abs(x, y) <= t {
			return true
		}
		if y, has := mp[id+1]; has && abs(x, y) <= t {
			return true
		}
		mp[id] = x
		if i >= k {
			delete(mp, getID(nums[i-k], t+1))
		}
	}
	return false
}
