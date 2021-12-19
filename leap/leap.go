// Package determines if a given year is a leap year
package leap

// Function with leap year logic
func IsLeapYear(y int) bool {
	return (y%4) == 0 && (y%100) != 0 || (y%400) == 0
}
