package reverse_string

func ReverseString(input string) (output string) {
	for _, lit := range input {
		output = string(lit) + output
	}
	return output
}
