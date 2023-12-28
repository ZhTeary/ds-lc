package search

import "sort"

func SuccessfulPairs(spells []int, potions []int, success int64) []int {
	result := make([]int, 0, len(spells))
	//对药水强度进行排序
	sort.Ints(potions)
	for i := range spells {
		//找到该咒语对应的药水最低强度
		least := 0
		if int(success)%spells[i] == 0 {
			least = int(success) / spells[i]
		} else {
			least = (int(success) / spells[i]) + 1
		}

		//在药水中找到最小的大于等于最低强度的药水，然后看后面还有几个
		ads := BinarySearch(potions, least)
		result = append(result, len(potions)-ads)
	}

	return result

	/*	超时
		result := make([]int, 0, len(spells))
		for i := range spells {
			count := 0
			for j := range potions {
				if potions[j]*spells[i] >= int(success) {
					count++
				}
			}
			result = append(result, count)
		}
		return result
	*/
}

func BinarySearch(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {

			//进行一个判断，如果前面还有相同的，那就继续找，如果前面没有相同的了，那就是第一个了
			if mid > 0 && nums[mid] == nums[mid-1] {
				right = mid - 1
				continue
			} else {
				return mid
			}

		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	//走到这儿就是left == right
	//返回的并不是该插入的位置，而是该插入的位置的前一个位置
	return left
}
