package sortlike

//插入排序
func InsertionSort(nums []int) []int {
	n := len(nums)

	for i := 1; i < n; i++ {
		//这里一定要有一个value变量记录nums[i] 不然nums[i]会被已排序区的最大值改变
		value := nums[i]

		for j := i - 1; j >= 0; j-- {
			if nums[j] > value { //这里也应比较value，因为nums[i]会被污染
				nums[j+1] = nums[j] //更改的元素应该都是num[j+1],因为j指向实际位置之前
			} else {
				nums[j+1] = value //这里应该是nums[j+1]=value 因为nums[i]已经被污染
				break             //在已经将元素插入合适位置后，就无需再遍历j了，因为前面已经有序
			}
		}

	}

	return nums
}
