package sortlike

func FindKthLargest(nums []int, k int) int {
	return quickSelect(nums, k)
}

func quickSelect(nums []int, k int) int {
	less := make([]int, 0, len(nums))
	equal := make([]int, 0, len(nums))
	greater := make([]int, 0, len(nums))
	pivot := nums[0]
	for i := range nums {
		if pivot > nums[i] {
			less = append(less, nums[i])
		} else if pivot < nums[i] {
			greater = append(greater, nums[i])
		} else {
			equal = append(equal, nums[i])
		}
	}
	if k <= len(greater) {
		return quickSelect(greater, k)
	}
	if len(nums)-k < len(nums) {
		return quickSelect(less, k-len(nums)+len(less))
	}
	return pivot
}
