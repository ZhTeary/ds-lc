package sortlike

//冒泡排序
func BubbleSort(nums []int) []int {

	for i := 0; i < len(nums); i++ { //每次都从第一个位置开始遍历

		//优化，当某次没有数据交换，可以退出循环
		exchange := false
		for j := 0; j < len(nums)-i-1; j++ { //从该位置遍历到以排序的位置为止 1.j从0开始 2. j<len()-i-1
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
				exchange = true
			}
		}

		//检验本次是否有冒泡操作
		if !exchange {
			break
		}
	}

	return nums

}
