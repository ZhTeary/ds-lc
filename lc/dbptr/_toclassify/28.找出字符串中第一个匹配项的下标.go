package doublePointer

func StrStr(haystack string, needle string) int {
	nee := 0

	//在saystack中循环遍历每个neesle
	for hay := 0; hay < len(haystack); hay++ {
		//当前匹配成功
		if haystack[hay] == needle[nee] {
			temp := hay
			for temp < len(haystack) && haystack[temp] == needle[nee] {
				nee++
				temp++
				if nee == len(needle) {
					return hay
				}
			}
			nee = 0
		}
	}
	return -1
}
