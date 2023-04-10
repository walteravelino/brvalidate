package extras

import (
	"github.com/walteravelino/brvalidate"
	"regexp"
	"strconv"
)

var (
	CepPattern = regexp.MustCompile("^\\d{5}-?\\d{3}$")
)

func CEP(value string, ufs ...brvalidate.FederativeUnit) bool {

	if !CepPattern.MatchString(value) {
		return false
	}

	h, _ := strconv.Atoi(value[0:3])

	if len(ufs) == 0 {
		return h >= 10
	}

	for _, uf := range ufs {
		if (uf == brvalidate.SP && h >= 10 && h <= 199) ||
			(uf == brvalidate.RJ && h >= 200 && h <= 289) ||
			(uf == brvalidate.ES && h >= 290 && h <= 299) ||
			(uf == brvalidate.MG && h >= 300 && h <= 399) ||
			(uf == brvalidate.BA && h >= 400 && h <= 489) ||
			(uf == brvalidate.SE && h >= 490 && h <= 499) ||
			(uf == brvalidate.PE && h >= 500 && h <= 569) ||
			(uf == brvalidate.AL && h >= 570 && h <= 579) ||
			(uf == brvalidate.PB && h >= 580 && h <= 589) ||
			(uf == brvalidate.RN && h >= 590 && h <= 599) ||
			(uf == brvalidate.CE && h >= 600 && h <= 639) ||
			(uf == brvalidate.PI && h >= 640 && h <= 649) ||
			(uf == brvalidate.MA && h >= 650 && h <= 659) ||
			(uf == brvalidate.PA && h >= 660 && h <= 688) ||
			(uf == brvalidate.AP && h == 689) ||
			(uf == brvalidate.AM && h >= 690 && h <= 692) ||
			(uf == brvalidate.RR && h == 693) ||
			(uf == brvalidate.AM && h >= 694 && h <= 698) ||
			(uf == brvalidate.AC && h == 699) ||
			(uf == brvalidate.DF && h >= 700 && h <= 727) ||
			(uf == brvalidate.GO && h >= 728 && h <= 729) ||
			(uf == brvalidate.DF && h >= 730 && h <= 736) ||
			(uf == brvalidate.GO && h >= 737 && h <= 767) ||
			(uf == brvalidate.RO && h >= 768 && h <= 769) ||
			(uf == brvalidate.TO && h >= 770 && h <= 779) ||
			(uf == brvalidate.MT && h >= 780 && h <= 788) ||
			(uf == brvalidate.MS && h >= 790 && h <= 799) ||
			(uf == brvalidate.PR && h >= 800 && h <= 879) ||
			(uf == brvalidate.SC && h >= 880 && h <= 899) ||
			(uf == brvalidate.RS && h >= 900 && h <= 999) {

			return true
		}
	}

	return false
}
