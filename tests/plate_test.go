package tests

import (
	brvalidate "github.com/walteravelino/br-validate"
	"testing"
)

func TestPlate(t *testing.T) {
	for i, item := range []struct {
		name     string
		expected bool
		value    string
	}{
		{"InvalidData", false, "3467875434578764345789654"},
		{"InvalidData", false, ""},
		{"InvalidData", false, "AAAAAAA"},
		{"ValidOldFormat", true, "AAA0000"},
		{"ValidOldFormat", true, "NFA-8324"},
		{"ValidNewFormat", true, "AAA0A00"},
		{"ValidNewFormat", true, "CAT4B31"},
	} {
		t.Run(testName(i, item.name), func(t *testing.T) {
			assertEqual(t, item.expected, brvalidate.Plate(item.value))
		})
	}
}

func TestNational(t *testing.T) {
	for i, item := range []struct {
		name     string
		expected bool
		value    string
	}{
		{"InvalidData", false, "3467875434578764345789654"},
		{"InvalidData", false, ""},
		{"InvalidData", false, "AAAAAAA"},
		{"InvalidNewFormat", false, "AAA0A00"},
		{"InvalidNewFormat", false, "ABC1D23"},
		{"ValidOldFormat", true, "AAA0000"},
		{"ValidOldFormat", true, "ABC-1234"},
	} {
		t.Run(testName(i, item.name), func(t *testing.T) {
			assertEqual(t, item.expected, brvalidate.National(item.value))
		})
	}
}

func TestMercosul(t *testing.T) {
	for i, item := range []struct {
		name     string
		expected bool
		value    string
	}{
		{"InvalidData", false, "3467875434578764345789654"},
		{"InvalidData", false, ""},
		{"InvalidData", false, "AAAAAAA"},
		{"InvalidOldFormat", false, "AAA0000"},
		{"InvalidOldFormat", false, "ABC-1234"},
		{"ValidNewFormat", true, "AAA0A00"},
		{"ValidNewFormat", true, "ABC1D23"},
	} {
		t.Run(testName(i, item.name), func(t *testing.T) {
			assertEqual(t, item.expected, brvalidate.Mercosul(item.value))
		})
	}
}
