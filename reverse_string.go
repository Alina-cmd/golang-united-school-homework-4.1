package reverse_string

func ReverseString(input string) (output string) {
	runes := []rune(input)
	newRunes := make([]rune, len(runes))
	count := 0
	for i := len(runes) - 1; i >= 0; i-- {
		newRunes[count] = runes[i]
		count++
	}
	output = string(newRunes)
	return output
}
