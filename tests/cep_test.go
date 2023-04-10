package tests

import (
	"github.com/walteravelino/brvalidate"
	"github.com/walteravelino/brvalidate/extras"
	"testing"
)

func TestCEP(t *testing.T) {
	for i, item := range []struct {
		name     string
		expected bool
		value    string
		uf       []domain.FederativeUnit
	}{
		{"InvalidData", false, "3467875434578764345789654", []domain.FederativeUnit{}},
		{"InvalidData", false, "", []domain.FederativeUnit{}},
		{"InvalidData", false, "AAAAAAAA", []domain.FederativeUnit{}},
		{"InvalidData", false, "00000-000", []domain.FederativeUnit{}},
		{"InvalidFormat", false, "10000 000", []domain.FederativeUnit{}},
		{"InvalidFormat", false, "25000.000", []domain.FederativeUnit{}},
		{"InvalidFederativeUnit", false, "10000-000", []domain.FederativeUnit{domain.RJ}},
		{"InvalidFederativeUnit", false, "25000-000", []domain.FederativeUnit{domain.ES}},
		{"InvalidFederativeUnit", false, "29500-000", []domain.FederativeUnit{domain.MG}},
		{"InvalidFederativeUnit", false, "35000-000", []domain.FederativeUnit{domain.BA}},
		{"InvalidFederativeUnit", false, "45000-000", []domain.FederativeUnit{domain.SE}},
		{"InvalidFederativeUnit", false, "49500-000", []domain.FederativeUnit{domain.PE}},
		{"InvalidFederativeUnit", false, "53000-000", []domain.FederativeUnit{domain.AL}},
		{"InvalidFederativeUnit", false, "57500-000", []domain.FederativeUnit{domain.PB}},
		{"InvalidFederativeUnit", false, "58500-000", []domain.FederativeUnit{domain.RN}},
		{"InvalidFederativeUnit", false, "59500-000", []domain.FederativeUnit{domain.CE}},
		{"InvalidFederativeUnit", false, "62000-000", []domain.FederativeUnit{domain.PI}},
		{"InvalidFederativeUnit", false, "64500-000", []domain.FederativeUnit{domain.MA}},
		{"InvalidFederativeUnit", false, "65500-000", []domain.FederativeUnit{domain.PA}},
		{"InvalidFederativeUnit", false, "67000-000", []domain.FederativeUnit{domain.AP}},
		{"InvalidFederativeUnit", false, "68900-000", []domain.FederativeUnit{domain.AM}},
		{"InvalidFederativeUnit", false, "69100-000", []domain.FederativeUnit{domain.RR}},
		{"InvalidFederativeUnit", false, "69300-000", []domain.FederativeUnit{domain.AM}},
		{"InvalidFederativeUnit", false, "69600-000", []domain.FederativeUnit{domain.AC}},
		{"InvalidFederativeUnit", false, "69900-000", []domain.FederativeUnit{domain.DF}},
		{"InvalidFederativeUnit", false, "71500-000", []domain.FederativeUnit{domain.GO}},
		{"InvalidFederativeUnit", false, "72800-000", []domain.FederativeUnit{domain.DF}},
		{"InvalidFederativeUnit", false, "73200-000", []domain.FederativeUnit{domain.GO}},
		{"InvalidFederativeUnit", false, "74500-000", []domain.FederativeUnit{domain.RO}},
		{"InvalidFederativeUnit", false, "76800-000", []domain.FederativeUnit{domain.TO}},
		{"InvalidFederativeUnit", false, "77500-000", []domain.FederativeUnit{domain.MT}},
		{"InvalidFederativeUnit", false, "78500-000", []domain.FederativeUnit{domain.MS}},
		{"InvalidFederativeUnit", false, "79500-000", []domain.FederativeUnit{domain.PR}},
		{"InvalidFederativeUnit", false, "84000-000", []domain.FederativeUnit{domain.SC}},
		{"InvalidFederativeUnit", false, "89000-000", []domain.FederativeUnit{domain.RS}},
		{"InvalidFederativeUnit", false, "95000-000", []domain.FederativeUnit{domain.SP}},
		{"InvalidFederativeUnit", false, "29500-000", []domain.FederativeUnit{domain.MT, domain.MS, domain.MG}},
		{"InvalidFederativeUnit", false, "00801-000", []domain.FederativeUnit{domain.SP}},
		{"InvalidFederativeUnit", false, "00801-000", []domain.FederativeUnit{}},
		{"Valid", true, "10000-000", []domain.FederativeUnit{domain.SP}},
		{"Valid", true, "25000-000", []domain.FederativeUnit{domain.RJ}},
		{"Valid", true, "29500-000", []domain.FederativeUnit{domain.ES}},
		{"Valid", true, "35000-000", []domain.FederativeUnit{domain.MG}},
		{"Valid", true, "45000-000", []domain.FederativeUnit{domain.BA}},
		{"Valid", true, "49500-000", []domain.FederativeUnit{domain.SE}},
		{"Valid", true, "53000-000", []domain.FederativeUnit{domain.PE}},
		{"Valid", true, "57500-000", []domain.FederativeUnit{domain.AL}},
		{"Valid", true, "58500-000", []domain.FederativeUnit{domain.PB}},
		{"Valid", true, "59500-000", []domain.FederativeUnit{domain.RN}},
		{"Valid", true, "62000-000", []domain.FederativeUnit{domain.CE}},
		{"Valid", true, "64500-000", []domain.FederativeUnit{domain.PI}},
		{"Valid", true, "65500-000", []domain.FederativeUnit{domain.MA}},
		{"Valid", true, "67000-000", []domain.FederativeUnit{domain.PA}},
		{"Valid", true, "68900-000", []domain.FederativeUnit{domain.AP}},
		{"Valid", true, "69100-000", []domain.FederativeUnit{domain.AM}},
		{"Valid", true, "69300-000", []domain.FederativeUnit{domain.RR}},
		{"Valid", true, "69600-000", []domain.FederativeUnit{domain.AM}},
		{"Valid", true, "69900-000", []domain.FederativeUnit{domain.AC}},
		{"Valid", true, "71500-000", []domain.FederativeUnit{domain.DF}},
		{"Valid", true, "72800-000", []domain.FederativeUnit{domain.GO}},
		{"Valid", true, "73200-000", []domain.FederativeUnit{domain.DF}},
		{"Valid", true, "74500-000", []domain.FederativeUnit{domain.GO}},
		{"Valid", true, "76800-000", []domain.FederativeUnit{domain.RO}},
		{"Valid", true, "77500-000", []domain.FederativeUnit{domain.TO}},
		{"Valid", true, "78500-000", []domain.FederativeUnit{domain.MT}},
		{"Valid", true, "79500-000", []domain.FederativeUnit{domain.MS}},
		{"Valid", true, "84000-000", []domain.FederativeUnit{domain.PR}},
		{"Valid", true, "89000-000", []domain.FederativeUnit{domain.SC}},
		{"Valid", true, "95000-000", []domain.FederativeUnit{domain.RS}},
		{"Valid", true, "10000-000", []domain.FederativeUnit{domain.SP, domain.SP, domain.SP}},
		{"Valid", true, "25000-000", []domain.FederativeUnit{}},
	} {
		t.Run(testName(i, item.name), func(t *testing.T) {
			assertEqual(t, item.expected, extras.CEP(item.value, item.uf...))
		})
	}
}

func TestCEPRegexp(t *testing.T) {
	for i, item := range []struct {
		name     string
		expected bool
		value    string
	}{
		{"Invalid", false, "3467875434578764345789654"},
		{"Invalid", false, ""},
		{"Invalid", false, "AAAAAAAA"},
		{"Invalid", false, "0-0000000"},
		{"Invalid", false, "00-000000"},
		{"Invalid", false, "000-00000"},
		{"Invalid", false, "0000-0000"},
		{"Invalid", false, "000000-00"},
		{"Invalid", false, "0000000-0"},
		{"Invalid", false, "0000x000"},
		{"Invalid", false, "0000x-000"},
		{"Invalid", false, "0000000x"},
		{"Invalid", false, "00000-00x"},
		{"Valid", true, "00000000"},
		{"Valid", true, "00000-000"},
		{"Valid", true, "12345678"},
		{"Valid", true, "87654-321"},
		{"Valid", true, "99999999"},
		{"Valid", true, "99999-999"},
	} {
		t.Run(testName(i, item.name), func(t *testing.T) {
			assertEqual(t, item.expected, extras.CepPattern.MatchString(item.value))
		})
	}
}
