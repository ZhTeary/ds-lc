package doublePointer

import (
	"math"
	"sort"
)

func ThreeSumClosest(nums []int, target int) int {
	sort.Ints(nums) //排序
	result := math.MaxInt

	//取绝对值
	abs := func(x int) int {
		if x < 0 {
			return -1 * x
		}
		return x
	}
	//更新差值
	update := func(current int) {
		if abs(current-target) < abs(result-target) {
			result = current
		}
	}

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] { // 这里i>0 是防止i-1越界
			continue
		}
		left := i + 1
		right := len(nums) - 1

		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == target {
				return sum
			}

			update(sum)

			if sum > target {
				right--
				for right > left && nums[right] == nums[right+1] {
					right--
				}

			} else if sum < target {
				left++
				for left < right && nums[left] == nums[left-1] {
					left++
				}
			}

		}

	}

	return result
}
