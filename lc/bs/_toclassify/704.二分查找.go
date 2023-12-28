package search

func BC(nums []int, target int) int {
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
	return left
}
