package tests

import (
	brvalidate "github.com/walteravelino/br-validate"
	"testing"
)

func TestCPF(t *testing.T) {
	for i, item := range []struct {
		name     string
		expected bool
		value    string
	}{
		{"InvalidData", false, "3467875434578764345789654"},
		{"InvalidData", false, ""},
		{"InvalidData", false, "AAAAAAAAAAA"},
		{"InvalidPattern", false, "000.000.000-00"},
		{"InvalidPattern", false, "111.111.111-11"},
		{"InvalidPattern", false, "222.222.222-22"},
		{"InvalidPattern", false, "333.333.333-33"},
		{"InvalidPattern", false, "444.444.444-44"},
		{"InvalidPattern", false, "555.555.555-55"},
		{"InvalidPattern", false, "666.666.666-66"},
		{"InvalidPattern", false, "777.777.777-77"},
		{"InvalidPattern", false, "888.888.888-88"},
		{"InvalidPattern", false, "999.999.999-99"},
		{"InvalidDigits", false, "248.438.034-08"},
		{"InvalidDigits", false, "099.075.865-06"},
		{"InvalidFormat", false, "248 438 034 80"},
		{"InvalidFormat", false, "099-075-865.60"},
		{"Valid", true, "172.507.700-04"},
		{"Valid", true, "054.276.100-90"},
	} {
		t.Run(testName(i, item.name), func(t *testing.T) {
			assertEqual(t, item.expected, brvalidate.CPF(item.value))
		})
	}
}
