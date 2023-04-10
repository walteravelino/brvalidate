package person

import (
	"regexp"
)

var (
	cnpjPattern = regexp.MustCompile("^\\d{2}\\.?\\d{3}\\.?\\d{3}/?(:?\\d{3}[1-9]|\\d{2}[1-9]\\d|\\d[1-9]\\d{2}|[1-9]\\d{3})-?\\d{2}$")
)

func CNPJ(value string) bool {

	const (
		size = 12
		pos  = 5
	)

	return DocumentValidate(value, cnpjPattern, size, pos)
}
