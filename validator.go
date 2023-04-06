package brvalidate

import "regexp"

func docValidate(doc string, pattern *regexp.Regexp, size int, position int) bool {

	if !pattern.MatchString(doc) {
		return false
	}

	cleansing(&doc)

	if sameDigits(doc) {
		return false
	}

	d := doc[:size]
	digit := dvCalc(d, position)

	d = d + digit
	digit = dvCalc(d, position+1)

	return doc == d+digit
}
