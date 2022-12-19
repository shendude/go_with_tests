package arrays

func Sum(arr []int) int {
	tot := 0
	for _, elem := range arr {
		tot += elem
	}
	return tot
}
