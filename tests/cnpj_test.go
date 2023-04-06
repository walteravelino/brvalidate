package tests

import (
	brvalidate "github.com/walteravelino/br-validate"
	"testing"
)

func TestCNPJ(t *testing.T) {
	for i, item := range []struct {
		name     string
		expected bool
		value    string
	}{
		{"InvalidData", false, "3467875434578764345789654"},
		{"InvalidData", false, ""},
		{"InvalidData", false, "AAAAAAAAAAAAAA"},
		{"InvalidPattern", false, "00.000.000/0000-00"},
		{"InvalidPattern", false, "11.111.111/1111-11"},
		{"InvalidPattern", false, "22.222.222/2222-22"},
		{"InvalidPattern", false, "33.333.333/3333-33"},
		{"InvalidPattern", false, "44.444.444/4444-44"},
		{"InvalidPattern", false, "55.555.555/5555-55"},
		{"InvalidPattern", false, "66.666.666/6666-66"},
		{"InvalidPattern", false, "77.777.777/7777-77"},
		{"InvalidPattern", false, "88.888.888/8888-88"},
		{"InvalidPattern", false, "99.999.999/9999-99"},
		{"InvalidDigits", false, "26.637.142/0001-85"},
		{"InvalidDigits", false, "74.221.325/0001-03"},
		{"InvalidFormat", false, "26-637-142.0001/58"},
		{"InvalidFormat", false, "74-221-325.0001/30"},
		{"Valid", true, "73.919.361/0001-00"},
		{"Valid", true, "90.522.791/0001-02"},
	} {
		t.Run(testName(i, item.name), func(t *testing.T) {
			assertEqual(t, item.expected, brvalidate.CNPJ(item.value))
		})
	}
}
