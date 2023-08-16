package exercise2

// Filter filters a list of integers based on the provided condition.
func Filter(nums []int, condition func(int) bool) []int {
	var result []int
	for _, num := range nums {
		if condition(num) {
			result = append(result, num)
		}
	}
	return result
}
