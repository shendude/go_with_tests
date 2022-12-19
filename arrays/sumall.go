package arrays

// SumAll returns the sums of multiple slices
// at the same time as a result slice.
func SumAll(arrs ...[]int) []int {
	// grab the length form the first element.
	var sums []int

	for _, elem := range arrs {
		sums = append(sums, Sum(elem))
	}
	return sums
}
