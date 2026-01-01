package iteration

import "strings"

const repeatCount = 5

/*
Takes a character, and returns it repeated loop number of times.
If given non-positive integer, defaults to 5 repeats.
*/
func Repeat(character string, loop int) string {
	var repeated strings.Builder

	if loop <= 0 {
		loop = repeatCount
	}

	for i := 0; i < loop; i++ {
		repeated.WriteString(character)
	}

	return repeated.String()
}
