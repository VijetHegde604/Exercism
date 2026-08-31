package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
	count := 0

	if n <= 0 {
		return 0, errors.New("Integer should be positive")
	}
	for n != 1 {
		if n % 2 == 0 {
			n = n / 2
		} else {
			n = 3 * n + 1
		}
		count += 1
	}
	return count, nil
	panic("Please implement the CollatzConjecture function")
}
