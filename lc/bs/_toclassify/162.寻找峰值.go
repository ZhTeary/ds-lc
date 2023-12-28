package search

import "math"

/*
O(logN)解决 ??	二分查找就是O(logN)
*/
func findPeakElement(nums []int) int {
	n := len(nums)

	get := func(i int) int {
		if i == -1 || i == n {
			return math.MaxInt64
		}
		return nums[i]
	}

	left, right := 0, n-1

	for { //这里不必left<= right
		mid := (left + right) / 2
		if get(mid-1) < get(mid) && get(mid) > get(mid+1) {
			return mid
		}
		if get(mid) < get(mid+1) {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
}
