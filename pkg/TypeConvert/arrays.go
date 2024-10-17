package typeconvert

import "strings"

func ArrayOfBytesToArrayOfStrings(input []byte) []string {
	str1 := string(input[:])
	return strings.Split(str1, "\n")
}
