package iteration

func Repeat(character string, numberOfRepetitions int8) string {
	var finalString string

	for i := 0; i < int(numberOfRepetitions); i++ {
		finalString += character
	}

	return finalString
}
