package ledger

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

type Entry struct {
	Date        string // "Y-m-d"
	Description string
	Change      int // in cents
}

// FormatLedger formats a slice of ledger entries into a readable string.
func FormatLedger(currency string, locale string, entries []Entry) (string, error) {
	// 1. Validate arguments immediately
	if currency != "USD" && currency != "EUR" {
		return "", errors.New("invalid currency")
	}
	if locale != "en-US" && locale != "nl-NL" {
		return "", errors.New("invalid locale")
	}

	// 2. Clone the entries slice safely so we don't mutate the caller's data
	entriesCopy := make([]Entry, len(entries))
	copy(entriesCopy, entries)

	// 3. Sort entries (Date -> Description -> Change)
	sort.Slice(entriesCopy, func(i, j int) bool {
		if entriesCopy[i].Date != entriesCopy[j].Date {
			return entriesCopy[i].Date < entriesCopy[j].Date
		}
		if entriesCopy[i].Description != entriesCopy[j].Description {
			return entriesCopy[i].Description < entriesCopy[j].Description
		}
		return entriesCopy[i].Change < entriesCopy[j].Change
	})

	var sb strings.Builder

	// 4. Write the header
	if locale == "nl-NL" {
		sb.WriteString(fmt.Sprintf("%-10s | %-25s | %-13s\n", "Datum", "Omschrijving", "Verandering"))
	} else {
		sb.WriteString(fmt.Sprintf("%-10s | %-25s | %-13s\n", "Date", "Description", "Change"))
	}

	// 5. Process and write each entry
	for _, entry := range entriesCopy {
		dateStr, err := formatDate(entry.Date, locale)
		if err != nil {
			return "", err
		}

		descStr := formatDescription(entry.Description)
		changeStr := formatCurrency(entry.Change, currency, locale)

		// %13s is right-aligned, %-25s is left-aligned
		sb.WriteString(fmt.Sprintf("%-10s | %-25s | %13s\n", dateStr, descStr, changeStr))
	}

	return sb.String(), nil
}

func formatDate(date, locale string) (string, error) {
	if len(date) != 10 || date[4] != '-' || date[7] != '-' {
		return "", errors.New("invalid date format")
	}

	y, m, d := date[0:4], date[5:7], date[8:10]
	if locale == "nl-NL" {
		return fmt.Sprintf("%s-%s-%s", d, m, y), nil
	}
	return fmt.Sprintf("%s/%s/%s", m, d, y), nil
}

func formatDescription(desc string) string {
	// Using []rune ensures we don't cut multi-byte characters in half
	runes := []rune(desc)
	if len(runes) > 25 {
		return string(runes[:22]) + "..."
	}
	return desc
}

func formatCurrency(change int, currency, locale string) string {
	negative := change < 0
	if negative {
		change = -change
	}

	// Guarantee at least 3 digits (e.g. 0 becomes "000") so we can safely separate cents
	centsStr := fmt.Sprintf("%03d", change)
	intPart := centsStr[:len(centsStr)-2]
	fracPart := centsStr[len(centsStr)-2:]

	// Separate integer part with thousands separators
	var parts []string
	for len(intPart) > 3 {
		parts = append(parts, intPart[len(intPart)-3:])
		intPart = intPart[:len(intPart)-3]
	}
	if len(intPart) > 0 {
		parts = append(parts, intPart)
	}

	// Reverse the slices since we built them backwards
	for i, j := 0, len(parts)-1; i < j; i, j = i+1, j-1 {
		parts[i], parts[j] = parts[j], parts[i]
	}

	// Join with correct locale separators
	var numStr string
	if locale == "nl-NL" {
		numStr = strings.Join(parts, ".") + "," + fracPart
	} else {
		numStr = strings.Join(parts, ",") + "." + fracPart
	}

	sym := "$"
	if currency == "EUR" {
		sym = "€"
	}

	// Return formatted output based on locale and sign
	if locale == "nl-NL" {
		if negative {
			return fmt.Sprintf("%s -%s ", sym, numStr)
		}
		return fmt.Sprintf("%s %s ", sym, numStr)
	}

	// en-US
	if negative {
		return fmt.Sprintf("(%s%s)", sym, numStr)
	}
	return fmt.Sprintf("%s%s ", sym, numStr)
}
