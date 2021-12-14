package raindrops

import (
	"strconv"
	"strings"
)

// - has 3 as a factor, add 'Pling' to the result.
// - has 5 as a factor, add 'Plang' to the result.
// - has 7 as a factor, add 'Plong' to the result.
// - does not have any of 3, 5, or 7 as a factor, the result should be the digits of the number.

func Convert(number int) string {
	var s strings.Builder
	if number%3 == 0 {
		s.WriteString("Pling")
	}
	if number%5 == 0 {
		s.WriteString("Plang")
	}
	if number%7 == 0 {
		s.WriteString("Plong")
	}
	result := s.String()
	if result == "" {
		return strconv.Itoa(number)
	}
	return result
}
