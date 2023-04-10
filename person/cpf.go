package person

import (
	"github.com/walteravelino/brvalidate"
	"regexp"
)

var (
	cpfPattern = regexp.MustCompile("^\\d{3}\\.?\\d{3}\\.?\\d{3}-?\\d{2}$")
)

func CPF(value string) bool {

	const (
		size = 9
		pos  = 10
	)

	return brvalidate.DocumentValidate(value, cpfPattern, size, pos)
}
