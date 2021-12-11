package hamming

import "errors"

func Distance(a, b string) (difference int, err error) {
	difference = 0

	if len(a) != len(b) {
		err = errors.New("DNA strands are not the same size")
	} else if len(a) > 0 {
		for i := 0; i < len(a); i++ {
			if string(a[i]) != string(b[i]) {
				difference += 1
			}
		}
	}
	return difference, err
}
