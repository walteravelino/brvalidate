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
		uf       []brvalidate.FederativeUnit
	}{
		{"InvalidData", false, "3467875434578764345789654", []brvalidate.FederativeUnit{}},
		{"InvalidData", false, "", []brvalidate.FederativeUnit{}},
		{"InvalidData", false, "AAAAAAAA", []brvalidate.FederativeUnit{}},
		{"InvalidData", false, "00000-000", []brvalidate.FederativeUnit{}},
		{"InvalidFormat", false, "10000 000", []brvalidate.FederativeUnit{}},
		{"InvalidFormat", false, "25000.000", []brvalidate.FederativeUnit{}},
		{"InvalidFederativeUnit", false, "10000-000", []brvalidate.FederativeUnit{brvalidate.RJ}},
		{"InvalidFederativeUnit", false, "25000-000", []brvalidate.FederativeUnit{brvalidate.ES}},
		{"InvalidFederativeUnit", false, "29500-000", []brvalidate.FederativeUnit{brvalidate.MG}},
		{"InvalidFederativeUnit", false, "35000-000", []brvalidate.FederativeUnit{brvalidate.BA}},
		{"InvalidFederativeUnit", false, "45000-000", []brvalidate.FederativeUnit{brvalidate.SE}},
		{"InvalidFederativeUnit", false, "49500-000", []brvalidate.FederativeUnit{brvalidate.PE}},
		{"InvalidFederativeUnit", false, "53000-000", []brvalidate.FederativeUnit{brvalidate.AL}},
		{"InvalidFederativeUnit", false, "57500-000", []brvalidate.FederativeUnit{brvalidate.PB}},
		{"InvalidFederativeUnit", false, "58500-000", []brvalidate.FederativeUnit{brvalidate.RN}},
		{"InvalidFederativeUnit", false, "59500-000", []brvalidate.FederativeUnit{brvalidate.CE}},
		{"InvalidFederativeUnit", false, "62000-000", []brvalidate.FederativeUnit{brvalidate.PI}},
		{"InvalidFederativeUnit", false, "64500-000", []brvalidate.FederativeUnit{brvalidate.MA}},
		{"InvalidFederativeUnit", false, "65500-000", []brvalidate.FederativeUnit{brvalidate.PA}},
		{"InvalidFederativeUnit", false, "67000-000", []brvalidate.FederativeUnit{brvalidate.AP}},
		{"InvalidFederativeUnit", false, "68900-000", []brvalidate.FederativeUnit{brvalidate.AM}},
		{"InvalidFederativeUnit", false, "69100-000", []brvalidate.FederativeUnit{brvalidate.RR}},
		{"InvalidFederativeUnit", false, "69300-000", []brvalidate.FederativeUnit{brvalidate.AM}},
		{"InvalidFederativeUnit", false, "69600-000", []brvalidate.FederativeUnit{brvalidate.AC}},
		{"InvalidFederativeUnit", false, "69900-000", []brvalidate.FederativeUnit{brvalidate.DF}},
		{"InvalidFederativeUnit", false, "71500-000", []brvalidate.FederativeUnit{brvalidate.GO}},
		{"InvalidFederativeUnit", false, "72800-000", []brvalidate.FederativeUnit{brvalidate.DF}},
		{"InvalidFederativeUnit", false, "73200-000", []brvalidate.FederativeUnit{brvalidate.GO}},
		{"InvalidFederativeUnit", false, "74500-000", []brvalidate.FederativeUnit{brvalidate.RO}},
		{"InvalidFederativeUnit", false, "76800-000", []brvalidate.FederativeUnit{brvalidate.TO}},
		{"InvalidFederativeUnit", false, "77500-000", []brvalidate.FederativeUnit{brvalidate.MT}},
		{"InvalidFederativeUnit", false, "78500-000", []brvalidate.FederativeUnit{brvalidate.MS}},
		{"InvalidFederativeUnit", false, "79500-000", []brvalidate.FederativeUnit{brvalidate.PR}},
		{"InvalidFederativeUnit", false, "84000-000", []brvalidate.FederativeUnit{brvalidate.SC}},
		{"InvalidFederativeUnit", false, "89000-000", []brvalidate.FederativeUnit{brvalidate.RS}},
		{"InvalidFederativeUnit", false, "95000-000", []brvalidate.FederativeUnit{brvalidate.SP}},
		{"InvalidFederativeUnit", false, "29500-000", []brvalidate.FederativeUnit{brvalidate.MT, brvalidate.MS, brvalidate.MG}},
		{"InvalidFederativeUnit", false, "00801-000", []brvalidate.FederativeUnit{brvalidate.SP}},
		{"InvalidFederativeUnit", false, "00801-000", []brvalidate.FederativeUnit{}},
		{"Valid", true, "10000-000", []brvalidate.FederativeUnit{brvalidate.SP}},
		{"Valid", true, "25000-000", []brvalidate.FederativeUnit{brvalidate.RJ}},
		{"Valid", true, "29500-000", []brvalidate.FederativeUnit{brvalidate.ES}},
		{"Valid", true, "35000-000", []brvalidate.FederativeUnit{brvalidate.MG}},
		{"Valid", true, "45000-000", []brvalidate.FederativeUnit{brvalidate.BA}},
		{"Valid", true, "49500-000", []brvalidate.FederativeUnit{brvalidate.SE}},
		{"Valid", true, "53000-000", []brvalidate.FederativeUnit{brvalidate.PE}},
		{"Valid", true, "57500-000", []brvalidate.FederativeUnit{brvalidate.AL}},
		{"Valid", true, "58500-000", []brvalidate.FederativeUnit{brvalidate.PB}},
		{"Valid", true, "59500-000", []brvalidate.FederativeUnit{brvalidate.RN}},
		{"Valid", true, "62000-000", []brvalidate.FederativeUnit{brvalidate.CE}},
		{"Valid", true, "64500-000", []brvalidate.FederativeUnit{brvalidate.PI}},
		{"Valid", true, "65500-000", []brvalidate.FederativeUnit{brvalidate.MA}},
		{"Valid", true, "67000-000", []brvalidate.FederativeUnit{brvalidate.PA}},
		{"Valid", true, "68900-000", []brvalidate.FederativeUnit{brvalidate.AP}},
		{"Valid", true, "69100-000", []brvalidate.FederativeUnit{brvalidate.AM}},
		{"Valid", true, "69300-000", []brvalidate.FederativeUnit{brvalidate.RR}},
		{"Valid", true, "69600-000", []brvalidate.FederativeUnit{brvalidate.AM}},
		{"Valid", true, "69900-000", []brvalidate.FederativeUnit{brvalidate.AC}},
		{"Valid", true, "71500-000", []brvalidate.FederativeUnit{brvalidate.DF}},
		{"Valid", true, "72800-000", []brvalidate.FederativeUnit{brvalidate.GO}},
		{"Valid", true, "73200-000", []brvalidate.FederativeUnit{brvalidate.DF}},
		{"Valid", true, "74500-000", []brvalidate.FederativeUnit{brvalidate.GO}},
		{"Valid", true, "76800-000", []brvalidate.FederativeUnit{brvalidate.RO}},
		{"Valid", true, "77500-000", []brvalidate.FederativeUnit{brvalidate.TO}},
		{"Valid", true, "78500-000", []brvalidate.FederativeUnit{brvalidate.MT}},
		{"Valid", true, "79500-000", []brvalidate.FederativeUnit{brvalidate.MS}},
		{"Valid", true, "84000-000", []brvalidate.FederativeUnit{brvalidate.PR}},
		{"Valid", true, "89000-000", []brvalidate.FederativeUnit{brvalidate.SC}},
		{"Valid", true, "95000-000", []brvalidate.FederativeUnit{brvalidate.RS}},
		{"Valid", true, "10000-000", []brvalidate.FederativeUnit{brvalidate.SP, brvalidate.SP, brvalidate.SP}},
		{"Valid", true, "25000-000", []brvalidate.FederativeUnit{}},
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
