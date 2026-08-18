package isbnverifier

import "strings"

func IsValidISBN(isbn string) bool {
	// Replace the dashes 
	isbnCleaned := strings.ReplaceAll(isbn, "-", "")

	N := len(isbnCleaned)
	aggregate := 0

	if N != 10 {
		return false
	}

	for i := range N - 1 {
		if isbnCleaned[i] < '0' || isbnCleaned[i] > '9' {
			return false
		}

		digit := int(isbnCleaned[i] - '0')
		aggregate += digit * (10 - i)
	}

	if isbnCleaned[N-1] == 'X' {
		aggregate += 10
	} else if isbnCleaned[N-1] >= '0' && isbnCleaned[N-1] <= '9' {
		aggregate += int(isbnCleaned[N-1] - '0')
	} else {
		return false
	}

	return aggregate%11 == 0
}
