package vehicle

import (
	"github.com/walteravelino/brvalidate/domain"
	"strconv"
)

var RenavanWeight = []int{3, 2, 9, 8, 7, 6, 5, 4, 3, 2}

func Renavan(value string) bool {
	if len(value) != 11 {
		return false
	}
	if !domain.SameDigits(value) {
		return false
	}

	var sum int
	for i, r := range value[:len(value)-1] {
		sum += domain.ToInt(r) * RenavanWeight[i]
	}

	digit := (sum * 10) % 11
	if digit == 10 {
		digit = 0
	}

	return string(value[len(value)-1]) == strconv.Itoa(digit)
}
