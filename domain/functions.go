package domain

import (
	"bytes"
	"regexp"
	"strconv"
	"unicode"
)

func DvCalculator(value string, pos int) string {
	var sum int

	for _, r := range value {
		sum += ToInt(r) * pos
		pos--

		if pos < 2 {
			pos = 9
		}
	}

	sum %= 11

	if sum < 2 {
		return "0"
	}

	return strconv.Itoa(11 - sum)
}

func Cleansing(value *string) {
	buf := bytes.NewBufferString("")

	for _, r := range *value {
		if unicode.IsDigit(r) {
			buf.WriteRune(r)
		}
	}

	*value = buf.String()
}

func SameDigits(value string) bool {
	data := value[0]

	for i := 1; i < len(value); i++ {
		if data != value[i] {
			return false
		}
	}

	return true
}

func ToInt(r rune) int {
	return int(r - '0')
}

func DocumentValidate(value string, pattern *regexp.Regexp, size int, pos int) bool {

	if !pattern.MatchString(value) {
		return false
	}

	Cleansing(&value)

	if SameDigits(value) {
		return false
	}

	d := value[:size]
	digit := DvCalculator(d, pos)

	d = d + digit
	digit = DvCalculator(d, pos+1)

	return value == d+digit
}

type FederativeUnit uint8

const (
	AC FederativeUnit = iota
	AL
	AP
	AM
	BA
	CE
	DF
	ES
	GO
	MA
	MT
	MS
	MG
	PA
	PB
	PR
	PE
	PI
	RJ
	RN
	RS
	RO
	RR
	SC
	SP
	SE
	TO
)
