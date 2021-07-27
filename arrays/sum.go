package sum

func Sum(nums []int) int {
	var result int
	for _, num := range nums {
		result += num
	}
	return result
}

func SumAll(numsToSum ...[]int) (sums []int) {
	for _, numbers := range numsToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}

func SumAllTails(numsToSum ...[]int) (sums []int) {
	for _, numbers := range numsToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return sums
}
