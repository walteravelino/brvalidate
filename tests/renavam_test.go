package tests

import (
	"github.com/walteravelino/brvalidate/domain/vehicle"
	"testing"
)

func TestRenavam(t *testing.T) {
	for i, item := range []struct {
		name     string
		expected bool
		value    string
	}{
		{"InvalidData", false, "3467875434578764345789654"},
		{"InvalidData", false, ""},
		{"InvalidData", false, "AAAAAAAAAAA"},
		{"InvalidDigit", false, "18800054176"},
		{"InvalidDigit", false, "70999776207"},
		{"InvalidDigit", false, "91889431490"},
		{"InvalidDigit", false, "02219943070"},
		{"Valid", true, "59937436844"},
		{"Valid", true, "02495642130"},
		{"Valid", true, "36835746560"},
		{"Valid", true, "45279164335"},
		{"Valid", true, "62948598212"},
		{"Valid", true, "01373045326"},
	} {
		t.Run(testName(i, item.name), func(t *testing.T) {
			assertEqual(t, item.expected, vehicle.Renavan(item.value))
		})
	}
}
