package armstrongnumbers

import "math"

func IsNumber(n int) bool {
	number := n
	numLen := 0

	for temp := number; temp > 0; temp /= 10 {
		numLen++
	}

	if number == 0 {
		numLen = 1
	}

	total := 0

	for number > 0 {
		digit := number % 10
		number /= 10

		total += int(math.Pow(float64(digit), float64(numLen)))
	}

	return n == total
}
