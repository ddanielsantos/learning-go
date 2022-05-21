package integers

func SumAllTails(slicesToSum ...[]int) []int {
	var sums []int

	for _, slice := range slicesToSum {
		// var lastItem = slice[len(slice)-1]
		// sums = append(sums, lastItem)

		if len(slice) == 0 {
			sums = append(sums, 0)
		} else {
			lastItem := slice[1:] // a slice that excludes the first item
			sums = append(sums, Sum(lastItem))
		}
	}

	return sums
}
