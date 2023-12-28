package sortlike

import "strings"

//简单

//	SortSentence func SortSentence(s string) string {
//		arr := strings.Split(s, " ")
//		sort.Slice(arr, func(i, j int) bool {
//			return arr[i][len(arr[i])-1] < arr[j][len(arr[j])-1]
//		})
//		for i, v := range arr {
//			arr[i] = v[:len(v)-1]
//		}
//
//		return strings.Join(arr, " ")
//	}
func SortSentence(s string) string {
	arr := strings.Split(s, " ")
	hmap := make(map[int]string)
	var sortarr []int

	for _, v := range arr {
		hmap[int(v[len(v)-1])] = v[:len(v)-1]
		sortarr = append(sortarr, int(v[len(v)-1]))
	}

	//插入
	for i := range sortarr {
		for j := i; j >= 1; j-- {
			if sortarr[j] < sortarr[j-1] {
				sortarr[j], sortarr[j-1] = sortarr[j-1], sortarr[j]
			}
		}
	}

	var payload []string

	for _, v := range sortarr {
		payload = append(payload, hmap[v])
	}
	return strings.Join(payload, " ")
}
