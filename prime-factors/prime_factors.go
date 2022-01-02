package prime

func Factors(n int64) []int64 {
	result := []int64{}
	num := uint64(n)

	for num%2 == 0 {
		num /= 2
		result = append(result, 2)
	}

	for i := uint64(3); i*i <= num; i += 2 {
		for num%i == 0 {
			num /= i
			result = append(result, int64(i))
		}
	}

	if num > 2 {
		result = append(result, int64(num))
	}
	return result
}
