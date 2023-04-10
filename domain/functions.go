package brvalidate

import (
	"bytes"
	"strconv"
	"unicode"
)

func dvCalc(doc string, pos int) string {
	var sum int

	for _, r := range doc {
		sum += toInt(r) * pos
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

func cleansing(doc *string) {
	buf := bytes.NewBufferString("")

	for _, r := range *doc {
		if unicode.IsDigit(r) {
			buf.WriteRune(r)
		}
	}

	*doc = buf.String()
}

func sameDigits(doc string) bool {
	data := doc[0]

	for i := 1; i < len(doc); i++ {
		if data != doc[i] {
			return false
		}
	}

	return true
}

func toInt(r rune) int {
	return int(r - '0')
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
