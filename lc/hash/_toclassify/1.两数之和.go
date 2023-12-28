package hash

func twoSum(nums []int, target int) []int {
	hashmap := make(map[int]int, len(nums))
	for i, v := range nums {
		caps := target - v
		index, ok := hashmap[caps]
		if ok {
			return []int{i, index}
		} else {
			hashmap[v] = i
		}
	}

	return nil
}
