package integers

// two arguments of the same type can be separeted with a comma
func Add(numbers []int) int {
	sum := 0

	// for i := 0; i < 5; i++ {
	// 	sum += numbers[i]
	// }

	for _, number := range numbers {
		sum += number
	}

	return sum
}
