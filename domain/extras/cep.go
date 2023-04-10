package extras

import (
	"github.com/walteravelino/brvalidate/domain"
	"regexp"
	"strconv"
)

var (
	CepPattern = regexp.MustCompile("^\\d{5}-?\\d{3}$")
)

func CEP(value string, ufs ...domain.FederativeUnit) bool {

	if !CepPattern.MatchString(value) {
		return false
	}

	h, _ := strconv.Atoi(value[0:3])

	if len(ufs) == 0 {
		return h >= 10
	}

	for _, uf := range ufs {
		if (uf == domain.SP && h >= 10 && h <= 199) ||
			(uf == domain.RJ && h >= 200 && h <= 289) ||
			(uf == domain.ES && h >= 290 && h <= 299) ||
			(uf == domain.MG && h >= 300 && h <= 399) ||
			(uf == domain.BA && h >= 400 && h <= 489) ||
			(uf == domain.SE && h >= 490 && h <= 499) ||
			(uf == domain.PE && h >= 500 && h <= 569) ||
			(uf == domain.AL && h >= 570 && h <= 579) ||
			(uf == domain.PB && h >= 580 && h <= 589) ||
			(uf == domain.RN && h >= 590 && h <= 599) ||
			(uf == domain.CE && h >= 600 && h <= 639) ||
			(uf == domain.PI && h >= 640 && h <= 649) ||
			(uf == domain.MA && h >= 650 && h <= 659) ||
			(uf == domain.PA && h >= 660 && h <= 688) ||
			(uf == domain.AP && h == 689) ||
			(uf == domain.AM && h >= 690 && h <= 692) ||
			(uf == domain.RR && h == 693) ||
			(uf == domain.AM && h >= 694 && h <= 698) ||
			(uf == domain.AC && h == 699) ||
			(uf == domain.DF && h >= 700 && h <= 727) ||
			(uf == domain.GO && h >= 728 && h <= 729) ||
			(uf == domain.DF && h >= 730 && h <= 736) ||
			(uf == domain.GO && h >= 737 && h <= 767) ||
			(uf == domain.RO && h >= 768 && h <= 769) ||
			(uf == domain.TO && h >= 770 && h <= 779) ||
			(uf == domain.MT && h >= 780 && h <= 788) ||
			(uf == domain.MS && h >= 790 && h <= 799) ||
			(uf == domain.PR && h >= 800 && h <= 879) ||
			(uf == domain.SC && h >= 880 && h <= 899) ||
			(uf == domain.RS && h >= 900 && h <= 999) {

			return true
		}
	}

	return false
}
