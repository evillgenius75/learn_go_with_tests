package sum

func Sum(nums []int) int {
	var result int
	for _, num := range nums {
		result += num
	}
	return result
}
