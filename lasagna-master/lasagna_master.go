package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, minPerLayer int) int {
	if minPerLayer == 0 {
		return len(layers) * 2
	}
	return len(layers) * minPerLayer
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodles int, sauce float64) {
	for _, v := range layers {
		if v == "sauce" {
			sauce += 0.2
		}
		if v == "noodles" {
			noodles += 50
		}
	}
	return
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsList, myList []string) []string {
	return append(myList, friendsList[len(friendsList)-1])
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, guests int) []float64 {
	var newRecipe []float64
	for _, v := range quantities {
		newRecipe = append(newRecipe, v*float64(guests)/float64(2))
	}
	return newRecipe
}
