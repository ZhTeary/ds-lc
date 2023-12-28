package dp

/*1137. 第N个泰波那契数 简单 2023-12-15 21:27-21:35 8min*/

func tribonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 || n == 2 {
		return 1
	}

	l1, l2, l3, sum := 0, 0, 1, 1
	for i := 3; i < n; i++ {
		l1 = l2
		l2 = l3
		l3 = sum
		sum = l1 + l2 + l3
	}
	return sum
}
