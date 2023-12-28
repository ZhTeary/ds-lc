package hash

import "fmt"

/*
	1657. 确定两个字符串是否接近 中等 2023/10/03 13:23 - 13:55 错了三次才成功

如果可以使用以下操作从一个字符串得到另一个字符串，则认为两个字符串 接近 ：

操作 1：交换任意两个 现有 字符。
例如，abcde -> aecdb
操作 2：将一个 现有 字符的每次出现转换为另一个 现有 字符，并对另一个字符执行相同的操作。
例如，aacabb -> bbcbaa（所有 a 转化为 b ，而所有的 b 转换为 a ）
你可以根据需要对任意一个字符串多次使用这两种操作。

给你两个字符串，word1 和 word2 。如果 word1 和 word2 接近 ，就返回 true ；否则，返回 false 。



示例 1：

输入：word1 = "abc", word2 = "bca"
输出：true
解释：2 次操作从 word1 获得 word2 。
执行操作 1："abc" -> "acb"
执行操作 1："acb" -> "bca"
示例 2：

输入：word1 = "a", word2 = "aa"
输出：false
解释：不管执行多少次操作，都无法从 word1 得到 word2 ，反之亦然。
示例 3：

输入：word1 = "cabbba", word2 = "abbccc"
输出：true
解释：3 次操作从 word1 获得 word2 。
执行操作 1："cabbba" -> "caabbb"
执行操作 2："caabbb" -> "baaccc"
执行操作 2："baaccc" -> "abbccc"
示例 4：

输入：word1 = "cabbba", word2 = "aabbss"
输出：false
解释：不管执行多少次操作，都无法从 word1 得到 word2 ，反之亦然。

*/

// CloseStrings 操作1 每个字符出现的次数相同 ; 操作2 两个字符出现的次数合相同
// 必须都得是现有字符，出现的综合次数相同
// 无序判断两个数组相等
func CloseStrings(word1 string, word2 string) bool {
	hash1 := make(map[string]int)
	hash2 := make(map[string]int)

	for _, v := range word1 {
		hash1[string(v)]++
	}
	for _, v := range word2 {
		hash2[string(v)]++
	}
	fmt.Println(hash1)
	fmt.Println(hash2)
	//对现有字符进行判断
	hashstr1 := make(map[string]bool)
	hashstr2 := make(map[string]bool)
	for i, _ := range hash1 {
		hashstr1[i] = true
	}
	for i, _ := range hash2 {
		hashstr2[i] = true
	}
	if equalmap(hashstr1, hashstr2) == false {
		return false
	}

	//对出现次数进行判断
	counts1 := make(map[int]int)
	counts2 := make(map[int]int)

	for _, i := range hash1 {
		counts1[i]++
	}
	for _, i := range hash2 {
		counts2[i]++
	}
	fmt.Println(counts1)
	fmt.Println(counts2)

	return equalmap(counts1, counts2)

}

func equalmap[M1, M2 ~map[K]V, K, V comparable](m1 M1, m2 M2) bool {
	if len(m1) != len(m2) {
		return false
	}
	for k, v1 := range m1 {
		if v2, ok := m2[k]; !ok || v1 != v2 {
			return false
		}
	}
	return true
}
