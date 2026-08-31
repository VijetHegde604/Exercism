package phonenumber

import (
	"fmt"
	"strings"
	"unicode"
)

// Number cleans and validates a phone number.
func Number(phoneNumber string) (string, error) {
	phoneNumber = strings.Map(func(r rune) rune {
		if unicode.IsDigit(r) {
			return r
		}
		return -1
	}, phoneNumber)

	// Remove the country code if present.
	if len(phoneNumber) == 11 {
		if phoneNumber[0] != '1' {
			return "", fmt.Errorf("Invalid country code")
		}
		phoneNumber = phoneNumber[1:]
	}

	// The number must contain exactly 10 digits.
	if len(phoneNumber) != 10 {
		return "", fmt.Errorf("number is not valid")
	}

	// Area and exchange codes must start with 2-9.
	if phoneNumber[0] < '2' || phoneNumber[0] > '9' ||
		phoneNumber[3] < '2' || phoneNumber[3] > '9' {
		return "", fmt.Errorf("Invalid phone number")
	}

	for i := range len(phoneNumber) {
		if i == 0 || i == 3 {
			continue
		}
		if phoneNumber[i] < '0' || phoneNumber[i] > '9' {
			return "", fmt.Errorf("Invalid phone number")
		}
	}

	return phoneNumber, nil
}

// AreaCode returns the first three digits of a valid phone number.
func AreaCode(phoneNumber string) (string, error) {
	phoneNumber, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}

	return phoneNumber[:3], nil
}

// Format returns a phone number in standard NANP format.
func Format(phoneNumber string) (string, error) {
	phoneNumber, err := Number(phoneNumber)
	if err != nil {
		return "", fmt.Errorf("Invalid number")
	}

	return fmt.Sprintf(
		"(%s) %s-%s",
		phoneNumber[:3],
		phoneNumber[3:6],
		phoneNumber[6:],
	), nil
}
