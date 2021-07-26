package sum

func Sum(nums [5]int) int {
	var result int
	for _, num := range nums {
		result += num
	}
	return result
}
