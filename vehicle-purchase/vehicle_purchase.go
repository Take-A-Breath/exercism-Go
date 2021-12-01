package purchase

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	switch {
	case kind == "car" || kind == "truck":
		return true
	}
	return false
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in dictionary order.
func ChooseVehicle(option1, option2 string) string {
	var choice string = ""
	if option1 < option2 {
		choice = option1 + " is clearly the better choice."
	} else {
		choice = option2 + " is clearly the better choice."
	}
	return choice
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	switch {
	case age < 3:
		return originalPrice * 0.80
	case age > 3 && age < 10:
		return originalPrice * 0.70
	case age >= 10:
		return originalPrice * 0.50
	default:
		return originalPrice
	}
}
