package luhn

// Valid determines if a given string identifier is a legitimate number
// according to the Luhn checksum formula.
func Valid(id string) bool {
	sum := 0
	digitCount := 0

	// Step 1: Iterate backwards through the string starting from the rightmost character.
	// This naturally sets up our right-to-left alternate doubling logic.
	for i := len(id) - 1; i >= 0; i-- {
		char := id[i]

		// Step 2: Dynamically ignore whitespace without changing the true index positions
		// of neighboring digits or illegal characters.
		if char == ' ' {
			continue
		}

		// Step 3: Convert the ASCII byte representation of the character to its
		// true numeric value using byte arithmetic ('0' is ASCII value 48).
		digit := int(char - '0')

		// Step 4: Stop immediately if any non-digit character (letters, punctuation) is found.
		if digit < 0 || digit > 9 {
			return false
		}

		// Step 5: Every second digit from the right must be doubled.
		// Since digitCount starts at 0, odd values (1, 3, 5...) capture every alternate index.
		if digitCount%2 == 1 {
			digit *= 2

			// Per Luhn rules, if doubling results in a two-digit number (10 to 18),
			// summing its digits is mathematically identical to subtracting 9.
			if digit > 9 {
				digit -= 9
			}
		}

		// Step 6: Update the running total and tracking metrics.
		sum += digit
		digitCount++
	}

	// Step 7: Enforce length constraints. The specification dictates that
	// strings containing fewer than two valid digits are automatically invalid.
	if digitCount <= 1 {
		return false
	}

	// Step 8: The identifier is valid only if the final sum is a multiple of 10.
	return sum%10 == 0
}
