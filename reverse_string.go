package reverse_string

import (
	"strings"
)

func ReverseString(input string) (output string) {
	sliceOfStrings := strings.Split(input, "\n")
	reversedSingleString := ""
	for i := 0; i < len(sliceOfStrings); i++ {
		if i > 0 {
			output += "\n"
		}
		reversedSingleString = ""
		for _, lit := range sliceOfStrings[i] {
			reversedSingleString = string(lit) + reversedSingleString
		}
		output += reversedSingleString
	}
	return output
}
