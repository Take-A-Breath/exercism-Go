package isogram

import "unicode"

func IsIsogram(word string) bool {
	m := make(map[rune]bool)

	for _, char := range word {
		upper := unicode.ToUpper(char)
		if !unicode.IsLetter(upper) {
			continue
		}
		_, is := m[upper]
		if is {
			return false
		} else {
			m[upper] = true
		}
	}
	return true
}
