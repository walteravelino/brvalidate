package tests

import (
	"github.com/walteravelino/brvalidate/domain/person"
	"testing"
)

func TestCNH(t *testing.T) {
	for i, item := range []struct {
		name     string
		expected bool
		value    string
	}{
		{"InvalidData", false, "3467875434578764345789654"},
		{"InvalidData", false, ""},
		{"InvalidData", false, "AAAAAAAAAAA"},
		{"InvalidDigit", false, "02102234243"},
		{"InvalidDigit", false, "02102234142"},
		{"InvalidDigit", false, "13798941353"},
		{"InvalidDigit", false, "00676003001"},
		{"InvalidDigit", false, "067600300-1"},
		{"InvalidDigit", false, "0067600300-1"},
		{"Valid", true, "34457036968"},
		{"Valid", true, "66171243725"},
		{"Valid", true, "94110258997"},
		{"Valid", true, "57428906095"},
		{"Valid", true, "36973144106"},
		{"Valid", true, "30582684883"},
	} {
		t.Run(testName(i, item.name), func(t *testing.T) {
			assertEqual(t, item.expected, person.CNH(item.value))
		})
	}
}
