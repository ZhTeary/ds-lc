package doublePointer

import (
	"sort"
)

/*
2023-10-29  第一次
2023-10-30 	第二次
2023-11-16  第三次
*/
func threeSum(nums []int) [][]int {
	//数组排序
	sort.Ints(nums)

	if len(nums) < 3 {
		return nil
	}

	//输出结果
	payload := make([][]int, 0, len(nums))

	//三重循环一定超时，三重->两重，第一层for，第二层while
	for i := range nums {

		if nums[i] > 0 {
			break
		}

		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		//定义左右指针
		left := i + 1
		right := len(nums) - 1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if sum == 0 {
				payload = append(payload, []int{nums[i], nums[left], nums[right]})
				//left为后一个不相等的
				left++
				for left > 0 && left < right && nums[left] == nums[left-1] {
					left++
				}
				//right为前一个不相等的
				right--
				for right < len(nums) && right > left && nums[right] == nums[right+1] {
					right--
				}
			} else if sum > 0 {
				right--
				for right < len(nums) && right > left && nums[right] == nums[right+1] {
					right--
				}
			} else {
				left++
				for left > 0 && left < right && nums[left] == nums[left-1] {
					left++
				}
			}
		}
	}
	return payload
}

// func ThreeSum(nums []int) [][]int {
// 	result := make([][]int, 0, len(nums))

// 	//将数组排序
// 	sort.Ints(nums)

// 	for i := 0; i < len(nums); i++ {
// 		if nums[i] > 0 {
// 			break
// 		}
// 		if i > 0 && nums[i] == nums[i-1] {
// 			continue
// 		}

// 		left := i + 1
// 		right := len(nums) - 1
// 		for left < right {
// 			sum := nums[i] + nums[left] + nums[right]
// 			if sum == 0 {
// 				result = append(result, []int{nums[i], nums[left], nums[right]})

// 				left++
// 				for left < right && nums[left] == nums[left-1] {
// 					left++
// 				}
// 				right--
// 				for right > left && nums[right] == nums[right+1] {
// 					right--
// 				}
// 			} else if sum > 0 {
// 				//右指针左移
// 				right--
// 				for right > left && nums[right] == nums[right+1] {
// 					right--
// 				}
// 			} else if sum < 0 {
// 				left++
// 				for left < right && nums[left] == nums[left-1] {
// 					left++
// 				}
// 			}
// 		}
// 	}

// 	return result
// }

// // 判断[][]中已经有[]了
// func Aisin(list [][]int, item []int) bool {
// 	flag := false
// 	for _, v := range list {
// 		if reflect.DeepEqual(v, item) {
// 			flag = true
// 		}
// 	}
// 	return flag
// }

/* 2023-10-29 第一次解题
func threeSum(nums []int) [][]int {
	result := make([][]int, 0, len(nums))
	nums = quicksort(nums)

	for first := range nums {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		//第三个数从最右开始
		third := len(nums) - 1
		target := -1 * nums[first]

		for second := first + 1; second < len(nums); second++ {
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			for second < third && nums[second]+nums[third] > target {
				third--
			}
			if second == third {
				break
			}
			if nums[second]+nums[third] == target {
				result = append(result, []int{nums[first], nums[second], nums[third]})
			}
		}

	}

	return result
}

func quicksort(nums []int) []int {
	if len(nums) < 2 {
		return nums

	}

	less := make([]int, 0, len(nums))
	equal := make([]int, 0, len(nums))
	greater := make([]int, 0, len(nums))

	pivot := nums[0]
	for i := range nums {
		if nums[i] < pivot {
			less = append(less, nums[i])
		} else if nums[i] == pivot {
			equal = append(equal, nums[i])
		} else {
			greater = append(greater, nums[i])
		}
	}
	return append(append(quicksort(less), equal...), quicksort(greater)...)
}
*/
