package arrays

func SumAllTails(arrs ...[]int) []int {
	var res []int
	for _, elem := range arrs {
		if (len(elem)) < 2 {
			res = append(res, 0)
		} else {
			res = append(res, Sum(elem[1:]))
		}
	}
	return res
}
