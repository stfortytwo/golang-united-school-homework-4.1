package reverse_string

import (
	"strings"
)

func ReverseString(input string) (output string) {
	for _, lit := range input {
		output = string(lit) + output
	}
	return output
}
