// common_elements.go

package exercise3

// CommonElements finds the common elements between two slices of integers.
func CommonElements(a, b []int) []int {
	common := make([]int, 0)
	seen := make(map[int]bool)

	for _, num := range a {
		seen[num] = true
	}

	for _, num := range b {
		if seen[num] {
			common = append(common, num)
		}
	}

	return common
}
