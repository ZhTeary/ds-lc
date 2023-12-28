package sortlike

func Quicksorter(nums []int) {
	Quicksort(nums, 0, len(nums)-1)
}

func Quicksort(nums []int, head int, tail int) {
	if head >= tail {
		return
	}

	pos := partition(nums, head, tail)
	Quicksort(nums, head, pos-1)
	Quicksort(nums, pos+1, tail)

}

func partition(nums []int, head int, tail int) int {

	i := head

	for j := head; j < tail; j++ {
		if nums[j] < nums[tail] {
			nums[j], nums[i] = nums[i], nums[j]
			i++
		}
	}

	nums[i], nums[tail] = nums[tail], nums[i]

	return i
}
