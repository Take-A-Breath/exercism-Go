package thefarm

import (
	"errors"
	"fmt"
)

// See types.go for the types defined for this exercise.

// TODO: Define the SillyNephewError type here.
type dimNephew struct {
	count int
}

func (e *dimNephew) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %d cows", e.count)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	if cows == 0 {
		return 0, errors.New("Division by zero")
	}
	if cows < 0 {
		return 0, &dimNephew{cows}
	}
	fodderAmount, err := weightFodder.FodderAmount()
	if err != nil {
		if err == ErrScaleMalfunction {
			fodderAmount = 2 * fodderAmount
		} else {
			return 0, err
		}
	}
	if fodderAmount < 0 {
		return 0, errors.New("Negative fodder")
	}
	return fodderAmount / float64(cows), nil
}
