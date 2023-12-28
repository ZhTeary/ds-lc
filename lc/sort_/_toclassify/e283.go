package sortlike

/* 简单 冒泡排序 O(n^2)*/
func MoveZero(nums []int) {
	//sort
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums)-1-i; j++ {
			if nums[j] == 0 {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}

/* 双指针 */
func MoveZeroes(nums []int) {
	//使用双指针方法
	//slower 指向位置 -- 0元素的位置
	//faster 指向目标 -- 非0元素的目标
	slower := 0
	for faster := 0; faster < len(nums); faster++ {
		if nums[slower] != 0 {
			slower++
		} else if nums[faster] != 0 {
			nums[faster], nums[slower] = nums[slower], nums[faster]
			slower++
		}

	}

}
