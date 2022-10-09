package main

func fairCandySwap(aliceSizes []int, bobSizes []int) []int {
	aSum := numsSum(aliceSizes)
	bSum := numsSum(bobSizes)
	sum := aSum + bSum
	am := initMap(aliceSizes, aSum, sum/2)
	bm := initMap(bobSizes, bSum, sum/2)

	for key, aVal := range am {

		if _, ok := bm[aVal]; ok {
			return []int{key, aVal}
		}
	}

	return []int{}
}

// key : a 交换的糖果数量
// value : a 需要的糖果数量
func initMap(nums []int, sum int, target int) map[int]int {
	m := make(map[int]int)

	for _, num := range nums {
		m[num] = target - (sum - num)
	}
	// fmt.Println (sum, target, m)
	return m
}

func numsSum(nums []int) int {
	sum := 0

	for _, num := range nums {
		sum += num
	}

	return sum
}
