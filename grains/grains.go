package grains

import "errors"

func Square(number int) (uint64, error) {
	switch {
	case number < 1 || number > 64:
		return 0, errors.New("invalid number")
	default:
		return 1 << (uint64(number) - 1), nil
	}
}

func Total() uint64 {
	return 1<<64 - 1
}
