package search

func MinEatingSpeed(piles []int, h int) int {

	maxi := 0
	//得到最大堆，结果在1-最大堆
	for _, pile := range piles {
		if pile > maxi {
			maxi = pile
		}
	}

	holeTime := func(piles []int, speed int) int {
		sum := 0
		for i := range piles {
			sum += (piles[i] + speed - 1) / speed
		}
		return sum
	}

	left := 1
	right := maxi
	for left <= right {
		mid := left + (right-left)/2

		if holeTime(piles, mid) > h {
			left = mid + 1
		} else if holeTime(piles, mid) <= h {
			right = mid - 1
		}
	}
	return left
}
