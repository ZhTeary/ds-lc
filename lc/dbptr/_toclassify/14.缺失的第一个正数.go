package doublePointer

func FirstMissingPositive(nums []int) int {

	//长度为n ，结果一定在[1,n+1]中间
	for i := range nums {
		if nums[i] <= 0 {
			nums[i] = len(nums) + 1
		}
	}

	//目前不为inf的数，就是数组中剩下的数了,数组中所有的数均为正数
	//用位置代替值，位置不同的第一个就是想要的结果
	for _, v := range nums {

		if abs(v) <= len(nums) {
			nums[abs(v)-1] = -abs(nums[abs(v)-1]) //右边也要-1的原因是如果数值和长度相等就溢出了
		}
	}

	//寻找第一个不为0
	for i := range nums {
		if nums[i] > 0 {
			return i + 1
		}
	}

	return len(nums) + 1
}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
