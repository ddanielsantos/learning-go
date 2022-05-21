package integers

// receive a list of slices and return a single containing the sums of them
func SumAll(numbersToSum ...[]int) []int {
	// numberOfSlices := len(numbersToSum)
	// sums = make([]int, numberOfSlices)
	var sums []int
	// for index, number := range numbersToSum {
	// 	sums[index] = Add(number)
	// }

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}

	return sums
}
