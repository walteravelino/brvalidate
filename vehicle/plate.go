package vehicle

import (
	"regexp"
)

var (
	NationalPattern = regexp.MustCompile("^[A-Z]{3}-?\\d{4}$")
	MercosulPattern = regexp.MustCompile("^[A-Z]{3}\\d[A-Z]\\d{2}$")
)

func National(value string) bool {
	return NationalPattern.MatchString(value)
}

func Mercosul(value string) bool {
	return MercosulPattern.MatchString(value)
}

func Plate(value string) bool {
	return National(value) || Mercosul(value)
}
