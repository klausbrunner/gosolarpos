// Package gosolarpos contains functions to find topocentric solar
// coordinates. That is, the position of the sun on the sky for a given
// latitude/longitude, time, and other parameters.
package gosolarpos

import "time"
import "math"

// Grena3 calculates topocentric solar position following algorithm number 3
// described in Grena, 'Five new algorithms for the computation of sun position
// from 2010 to 2110', Solar Energy 86 (2012) pp. 1323-1337.
// deltaT is the difference between universal time and terrestrial time, in seconds.
// pressure is air pressure in hectopascal, used for refraction correction (if unsure, use 1000).
// temperature is air temperature in degrees Celsius, used for refraction correction.
func Grena3(date time.Time, latitude float64, longitude float64, deltaT float64, pressure float64, temperature float64) (azimuthDegrees, zenithDegrees float64) {
	t := calcT(date)

	tE := t + 1.1574e-5*deltaT
	omegaAtE := 0.0172019715 * tE

	lambda := -1.388803 + 1.720279216e-2*tE + 3.3366e-2*math.Sin(omegaAtE-0.06172) + 3.53e-4*math.Sin(2.0*omegaAtE-0.1163)

	epsilon := 4.089567e-1 - 6.19e-9*tE

	sLambda := math.Sin(lambda)
	cLambda := math.Cos(lambda)
	sEpsilon := math.Sin(epsilon)
	cEpsilon := math.Sqrt(1.0 - sEpsilon*sEpsilon)

	alpha := math.Atan2(sLambda*cEpsilon, cLambda)
	if alpha < 0 {
		alpha = alpha + 2*math.Pi
	}

	delta := math.Asin(sLambda * sEpsilon)

	h := 1.7528311 + 6.300388099*t + toRad(longitude) - alpha
	h = math.Mod((h+math.Pi), (2*math.Pi)) - math.Pi
	if h < -math.Pi {
		h = h + 2*math.Pi
	}

	sPhi := math.Sin(toRad(latitude))
	cPhi := math.Sqrt((1 - sPhi*sPhi))
	sDelta := math.Sin(delta)
	cDelta := math.Sqrt(1 - sDelta*sDelta)
	sH := math.Sin(h)
	cH := math.Cos(h)

	sEpsilon0 := sPhi*sDelta + cPhi*cDelta*cH
	eP := math.Asin(sEpsilon0) - 4.26e-5*math.Sqrt(1.0-sEpsilon0*sEpsilon0)
	gamma := math.Atan2(sH, cH*sPhi-sDelta*cPhi/cDelta)

	deltaRe := 0.0
	if eP > 0 && pressure > 0.0 && pressure < 3000.0 && temperature > -273 && temperature < 273 {
		deltaRe = (0.08422 * (pressure / 1000)) / ((273.0 + temperature) * math.Tan(eP+0.003138/(eP+0.08919)))
	}

	z := math.Pi/2 - eP - deltaRe

	azimuthDegrees = convertAzimuth(gamma)
	zenithDegrees = toDeg(z)
	return
}

func convertAzimuth(radFromSouth float64) (degFromNorth float64) {
	degFromNorth = math.Mod(toDeg(radFromSouth+math.Pi), 360.0)
	return
}

func toRad(degs float64) float64 {
	return degs / 180.0 * math.Pi
}

func toDeg(rads float64) float64 {
	return rads * 180.0 / math.Pi
}

func calcT(date time.Time) float64 {
	utc := date.UTC()

	m := float64(utc.Month())
	y := float64(utc.Year())
	d := float64(utc.Day())
	h := float64(utc.Hour()) +
		float64(utc.Minute())/60.0 +
		float64(utc.Second())/(60.0*60.0)

	if m <= 2 {
		m = m + 12
		y = y - 1
	}

	return float64(int(365.25*(y-2000))) + float64(int(30.6001*(m+1))) - float64(int(0.01*y)) + d + 0.0416667*h - 21958
}
