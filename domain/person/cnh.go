package person

import "github.com/walteravelino/brvalidate/domain"

func CNH(value string) bool {
	if len(value) != 11 {
		return false
	}
	if !domain.SameDigits(value) {
		return false
	}

	sum := 0
	weight := 9

	for _, r := range value[:len(value)-2] {
		sum += domain.ToInt(r) * weight
		weight--
	}

	base := 0
	firstDigit := sum % 11

	if firstDigit == 10 {
		base = -2
	}
	if firstDigit > 9 {
		firstDigit = 0
	}

	sum = 0
	weight = 1

	for _, r := range value[:len(value)-2] {
		sum += domain.ToInt(r) * weight
		weight++
	}

	var secondDigit int

	if (sum%11)+base < 0 {
		secondDigit = 11 + (sum % 11) + base
	}
	if (sum%11)+base >= 0 {
		secondDigit = (sum % 11) + base
	}
	if secondDigit > 9 {
		secondDigit = 0
	}

	return domain.ToInt(rune(value[len(value)-2])) == firstDigit &&
		domain.ToInt(rune(value[len(value)-1])) == secondDigit
}
