package doublePointer

import "math"

func nextPermutation(nums []int) {
	//15:15 - 16:15
	head := 0
	tail := 1
	// 最后出现头<=尾的位置
	last := -1

	for tail < len(nums) {
		if nums[head] < nums[last] {
			last = head
		}
		head++
		tail++
	}
	//如果已经是最大了
	if last == -1 {
		fullTurn(nums)
	} else if last == len(nums)-2 { //如果最后一个升序在结尾,只调换结尾顺序
		nums[head], nums[tail] = nums[tail], nums[head]
	} else { //如果最后一个升序在中间
		nxt := findNext(nums, last)
		//交换顺序
		nums[nxt], nums[last] = nums[last], nums[nxt]
		//重新排序
		sort_here(nums[last+1:])
	}

}
func fullTurn(nums []int) {
	head := 0
	tail := len(nums) - 1
	for head < tail {
		nums[head], nums[tail] = nums[tail], nums[head]
		head++
		tail--
	}
}

func findNext(nums []int, idx int) int {
	caps := math.MaxInt
	target := -1
	for i := idx + 1; i < len(nums); i++ {
		if nums[i] > nums[idx] {
			if nums[i]-nums[idx] < caps {
				target = i
				caps = nums[i] - nums[idx]
			}
		}
	}
	return target
}

func sort_here(nums []int) {
	//选择-原地
	for i := 0; i < len(nums)-1; i++ {
		mini := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[mini] {
				mini = j
			}
		}
		nums[mini], nums[i] = nums[i], nums[mini]
	}
}
