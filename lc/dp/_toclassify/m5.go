package dp

import "fmt"

/*
	5.最长回文子串 medium 2023/12/18 5:40~
*/

//中心扩展

//动态规划，自己没写上来， 看了思路之后自己写的
func LongestPalindrome(s string) string {
	n := len(s)

	//二维数组初始化
	dp := make([][]bool, n) //坑0. 此问题dp数组中存的是状态，而不是结果值，应注意
	for i := range dp {     //坑1. 二维数组的初始化，要在一个循环中单独初始化第二维
		dp[i] = make([]bool, n)
	}

	//dp数组赋值
	for i := 0; i < n; i++ { // i>=j 全为true
		dp[i][i] = true
		for j := 0; j <= i; j++ { //坑2. 既然dp中存状态，就要想要各种情况下的状态，并提早准备
			dp[i][j] = true
		}
	}

	//遍历s
	for i := n - 2; i >= 0; i-- { //坑3. 此种情况一定要从后往前遍历，确保在计算父问题时，子问题已经处于解决完的状态
		for j := n - 1; j > i; j-- { //坑4. 在遍历j时，一定要注意j的临界条件，提前赋值好的状态不要再遍历了
			dp[i][j] = dp[i+1][j-1] && s[i] == s[j]
		}
	}

	//查找结果
	begin := 0
	end := 0
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if dp[i][j] == true && j-i > (end-begin) {
				begin = i
				end = j
			}
		}
	}
	fmt.Println(dp)

	return string(s[begin : end+1]) //坑5. 取切片时是左闭右开的
}
