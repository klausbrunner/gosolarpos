package gosolarpos

import (
	"math"
	"time"
)

// EstimateDeltaT returns an inter- or extrapolated estimate of Delta T for the
// given time. This is based on Espenak and Meeus, "Five Millennium Canon of
// Solar Eclipses: -1999 to +3000" (NASA/TP-2006-214141).
func EstimateDeltaT(forDate time.Time) (deltaTSeconds float64) {
	year := decimalYear(forDate)

	if year < -500 {
		u := (year - 1820) / 100
		deltaTSeconds = -20 + 32*math.Pow(u, 2)
	} else if year < 500 {
		u := year / 100
		deltaTSeconds = 10583.6 - 1014.41*u + 33.78311*math.Pow(u, 2) -
			5.952053*math.Pow(u, 3) - 0.1798452*math.Pow(u, 4) + 0.022174192*math.Pow(u, 5) +
			0.0090316521*math.Pow(u, 6)
	} else if year < 1600 {
		u := (year - 1000) / 100
		deltaTSeconds = 1574.2 - 556.01*u + 71.23472*math.Pow(u, 2) +
			0.319781*math.Pow(u, 3) - 0.8503463*math.Pow(u, 4) - 0.005050998*math.Pow(u, 5) +
			0.0083572073*math.Pow(u, 6)
	} else if year < 1700 {
		t := year - 1600
		deltaTSeconds = 120 - 0.9808*t - 0.01532*math.Pow(t, 2) + math.Pow(t, 3)/7129
	} else if year < 1800 {
		t := year - 1700
		deltaTSeconds = 8.83 + 0.1603*t - 0.0059285*math.Pow(t, 2) +
			0.00013336*math.Pow(t, 3) - math.Pow(t, 4)/1174000
	} else if year < 1860 {
		t := year - 1800
		deltaTSeconds = 13.72 - 0.332447*t + 0.0068612*math.Pow(t, 2) +
			0.0041116*math.Pow(t, 3) - 0.00037436*math.Pow(t, 4) +
			0.0000121272*math.Pow(t, 5) - 0.0000001699*math.Pow(t, 6) + 0.000000000875*math.Pow(t, 7)
	} else if year < 1900 {
		t := year - 1860
		deltaTSeconds = 7.62 + 0.5737*t - 0.251754*math.Pow(t, 2) +
			0.01680668*math.Pow(t, 3) - 0.0004473624*math.Pow(t, 4) + math.Pow(t, 5)/233174
	} else if year < 1920 {
		t := year - 1900
		deltaTSeconds = -2.79 + 1.494119*t - 0.0598939*math.Pow(t, 2) +
			0.0061966*math.Pow(t, 3) - 0.000197*math.Pow(t, 4)
	} else if year < 1941 {
		t := year - 1920
		deltaTSeconds = 21.20 + 0.84493*t - 0.076100*math.Pow(t, 2) +
			0.0020936*math.Pow(t, 3)
	} else if year < 1961 {
		t := year - 1950
		deltaTSeconds = 29.07 + 0.407*t - math.Pow(t, 2)/233 +
			math.Pow(t, 3)/2547
	} else if year < 1986 {
		t := year - 1975
		deltaTSeconds = 45.45 + 1.067*t - math.Pow(t, 2)/260 -
			math.Pow(t, 3)/718
	} else if year < 2005 {
		t := year - 2000
		deltaTSeconds = 63.86 + 0.3345*t - 0.060374*math.Pow(t, 2) +
			0.0017275*math.Pow(t, 3) + 0.000651814*math.Pow(t, 4) + 0.00002373599*math.Pow(t, 5)
	} else if year < 2050 {
		t := year - 2000
		deltaTSeconds = 62.92 + 0.32217*t + 0.005589*math.Pow(t, 2)
	} else if year < 2150 {
		deltaTSeconds = -20 + 32*math.Pow(((year-1820)/100), 2) - 0.5628*(2150-year)
	} else {
		u := (year - 1820) / 100
		deltaTSeconds = -20 + 32*math.Pow(u, 2)
	}
	return
}

func decimalYear(forDate time.Time) float64 {
	rawYear := float64(forDate.Year())

	return rawYear + (float64(forDate.Month())-0.5)/12.0
}
