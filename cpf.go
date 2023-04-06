package brvalidate

import (
	"regexp"
)

var (
	cpfPattern = regexp.MustCompile("^\\d{3}\\.?\\d{3}\\.?\\d{3}-?\\d{2}$")
)

func cpfValidate(doc string) bool {

	const (
		size = 9
		pos  = 10
	)

	return docValidate(doc, cpfPattern, size, pos)
}
