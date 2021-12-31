package luhn

import (
	"strings"
	"unicode"
)

func Valid(id string) bool {
	counter := 0

	// Spaces are allowed in the input, but they should be stripped before checking
	id = strings.ReplaceAll(id, " ", "")
	parity := len(id) % 2

	// other non-digit characters are disallowed.
	for pos, chr := range id {
		if !unicode.IsDigit(chr) {
			return false
		}
		num := int(chr - '0')
		// double every second digit, starting from the right.
		if (pos % 2) == parity {
			num *= 2
			// If doubling the number results in a number greater than 9 then subtract 9 from the product.
			if num > 9 {
				num -= 9
			}
		}
		counter += num
	}
	// If the sum is evenly divisible by 10, then the number is valid.
	return (counter%10) == 0 && len(id) > 1
}
