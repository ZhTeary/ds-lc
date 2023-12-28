package hash

//2023/10/06 21:46-22:04 中等 18分钟

func letterCombinations(digits string) []string {
	hashmap := make(map[string][]string, 8)
	hashmap["2"] = []string{"a", "b", "c"}
	hashmap["3"] = []string{"d", "e", "f"}
	hashmap["4"] = []string{"g", "h", "i"}
	hashmap["5"] = []string{"j", "k", "l"}
	hashmap["6"] = []string{"m", "n", "o"}
	hashmap["7"] = []string{"p", "q", "r", "s"}
	hashmap["8"] = []string{"t", "u", "v"}
	hashmap["9"] = []string{"w", "x", "y", "z"}

	var payload []string
	if len(digits) == 0 {
		return payload
	} else {
		payload = hashmap[string(digits[0])]
	}
	//string
	for i := 1; i < len(digits); i++ {
		payload = dicar(payload, hashmap[string(digits[i])])

	}

	return payload

}

func dicar(x, y []string) []string {
	var ans []string
	for _, v := range x {
		for _, u := range y {
			ans = append(ans, v+u)
		}
	}
	return ans
}
