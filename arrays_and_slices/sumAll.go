package arrays_and_slices

func SumAll(numbersToSum ...[]int) []int {
	sums := []int{}

	for _, numbers := range numbersToSum {
		 tail := numbers[1:] 
		sums = append(sums, Sum(tail))
	}

	return sums
}
