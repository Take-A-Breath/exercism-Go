package cars

import "fmt"

// SuccessRate is used to calculate the ratio of an item being created without
// error for a given speed
// base rate = 221
func SuccessRate(speed int) float64 {
	var rate int = 221
	if speed == 0 {
		fmt.Println(float64(rate) * 0)
	} else if speed >= 1 && speed <= 4 {
		fmt.Println(float64(1 * rate))
	} else if speed >= 5 && speed <= 8 {
		fmt.Println(float64(rate) * 0.90)
	} else if speed >= 9 && speed <= 10 {
		fmt.Println(float64(rate) * 0.77)
	} else {
		fmt.Println(rate, speed)
	}
}

// CalculateProductionRatePerHour for the assembly line, taking into account
// its success rate
func CalculateProductionRatePerHour(speed int) float64 {
	panic("CalculateProductionRatePerHour not implemented")
}

// CalculateProductionRatePerMinute describes how many working items are
// produced by the assembly line every minute
func CalculateProductionRatePerMinute(speed int) int {
	panic("CalculateProductionRatePerMinute not implemented")
}

// CalculateLimitedProductionRatePerHour describes how many working items are
// produced per hour with an upper limit on how many can be produced per hour
func CalculateLimitedProductionRatePerHour(speed int, limit float64) float64 {
	panic("CalculateLimitedProductionRatePerHour not implemented")
}
