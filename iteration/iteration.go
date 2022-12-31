package iteration

func Repeat(character string, numberOfRepeats int) string {
	var repeated string
	for i := 0; i < numberOfRepeats; i++ {
		repeated += character
	}
	return repeated
}
