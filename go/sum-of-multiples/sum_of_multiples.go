package sumofmultiples

func SumMultiples(limit int, divisors ...int) int {
	multiples := make(map[int]bool)

	for _, divisor := range divisors {
		if divisor <= 0 {
			continue
		}
		for multiple := divisor; multiple < limit; multiple += divisor {
			multiples[multiple] = true
		}
	}

	sum := 0
	for multiple := range multiples {
		sum += multiple
	}
	return sum
}
