package slipWindow

func lengthOfLongestSubstring(s string) int {

	if len(s) == 0 {
		return 0
	}

	hashmap := make(map[byte]int, len(s))
	l, r, maxi := 0, 0, 0

	for ; r < len(s); r++ {
		if _, ok := hashmap[s[r]]; !ok {
			hashmap[s[r]] = r
		} else {
			//上次出现在本子串中，即出现重复,更新位置
			if l <= hashmap[s[r]] {
				l = hashmap[s[r]] + 1
			}
			hashmap[s[r]] = r
		}
		//每次循环都检查最大值
		if r-l+1 > maxi {
			maxi = r - l + 1
		}
	}
	return maxi
}
