package brvalidate

import (
	"regexp"
)

var (
	cpfPattern = regexp.MustCompile("^\\d{3}\\.?\\d{3}\\.?\\d{3}-?\\d{2}$")
)

func CPF(doc string) bool {

	const (
		size = 9
		pos  = 10
	)

	return DocValidate(doc, cpfPattern, size, pos)
}
