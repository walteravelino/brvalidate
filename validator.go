package brvalidate

import "regexp"

func DocValidate(doc string, pattern *regexp.Regexp, size int, pos int) bool {

	if !pattern.MatchString(doc) {
		return false
	}

	cleansing(&doc)

	if sameDigits(doc) {
		return false
	}

	d := doc[:size]
	digit := dvCalc(d, pos)

	d = d + digit
	digit = dvCalc(d, pos+1)

	return doc == d+digit
}
