package hamming
import "errors"
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("Only works with strings of same length")
	}
	count := 0
	for i := 0; i < len(a); i++ {
		if b[i] != a[i] {
			count += 1
		}
	}
	return count, nil
	panic("Implement the Distance function")
}
