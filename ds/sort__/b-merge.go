package sortlike

import (
	"fmt"
)

func Merger(nums []int) {
	MergeSort(nums, 0, len(nums)-1)
}

func MergeSort(nums []int, head int, tail int) {
	fmt.Println("本次归的是nums[", head, ",", tail, "]")
	if head == tail {
		return
	}
	//归并排序中的mid才是灵魂，归并排序归并的不是分开的数组，而是mid下标！
	mid := head + ((tail - head) / 2) //a+((b-a)/2)的写法
	fmt.Println("mid", mid)

	MergeSort(nums, head, mid)
	MergeSort(nums, mid+1, tail)
	Merge(nums, head, mid, tail)

}

func Merge(nums []int, head int, mid int, tail int) {
	fmt.Println("== 本次并的是nums[", head, ",", tail, "] ==", nums)

	//新建temp承接原数组进行排序，导致归并排序的空间复杂度为O(n)
	temp := make([]int, 0, tail-head+1)

	i := head
	j := mid + 1
	for i <= mid && j <= tail {
		if nums[i] <= nums[j] {
			temp = append(temp, nums[i])
			i++
		} else if nums[i] > nums[j] {
			temp = append(temp, nums[j])
			j++
		}
	}
	if i == mid+1 && j <= tail { //<= 很重要
		temp = append(temp, nums[j:tail+1]...)
	} else if j == tail+1 && i <= mid { //这里应该是小于等于mid
		temp = append(temp, nums[i:mid+1]...)
	}
	fmt.Println(temp)
	copy(nums[head:tail+1], temp)

}

// func Merge_Sort(nums []int, head int, tail int) {
// 	if head == tail {
// 		return
// 	}
// 	//这句话才是灵魂！
// 	mid := head + ((tail - head) / 2) //归并排序归并的不是分开的数组，而是mid下标！

// 	MergeSort(nums, head, mid)
// 	MergeSort(nums, mid+1, tail)
// 	Merge(nums, head, tail)
// }
