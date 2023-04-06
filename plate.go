package brvalidate

import (
	"regexp"
)

var (
	NationalPattern = regexp.MustCompile("^[A-Z]{3}-?\\d{4}$")
	MercosulPattern = regexp.MustCompile("^[A-Z]{3}\\d[A-Z]\\d{2}$")
)

func Plate(doc string) bool {
	return National(doc) || Mercosul(doc)
}

func National(doc string) bool {
	return NationalPattern.MatchString(doc)
}

func Mercosul(doc string) bool {
	return MercosulPattern.MatchString(doc)
}
