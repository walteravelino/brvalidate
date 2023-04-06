package brvalidate

func CNH(doc string) bool {
	if len(doc) != 11 {
		return false
	}
	if !sameDigits(doc) {
		return false
	}

	sum := 0
	acc := 9

	for _, r := range doc[:len(doc)-2] {
		sum += toInt(r) * acc
		acc--
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
	acc = 1

	for _, r := range doc[:len(doc)-2] {
		sum += toInt(r) * acc
		acc++
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

	return toInt(rune(doc[len(doc)-2])) == firstDigit &&
		toInt(rune(doc[len(doc)-1])) == secondDigit
}
