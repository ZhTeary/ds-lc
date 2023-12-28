package sortlike

//选择排序
func SelectionSort(nums []int) []int { //选择一个元素到合适位置

	for i := 0; i < len(nums); i++ {
		mini := i
		for j := i; j < len(nums); j++ {
			if nums[j] < nums[mini] { //是个坑，要找到最小的mini，所以nums[j]每次都要和nums[mini]比较
				mini = j
			}
		}
		nums[i], nums[mini] = nums[mini], nums[i]
	}

	return nums

}
