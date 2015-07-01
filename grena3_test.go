package gosolarpos

import (
	"fmt"
	"math"
	"testing"
	"time"
)

func ExampleGrena3() {
	date := time.Date(2015, 6, 23, 10, 30, 12, 0, time.UTC)

	azimuth, zenithAngle := Grena3(date,
		52.521667, // latitude (degrees)
		13.413333, // longitude (degrees)
		65,        // Delta T (s)
		1000,      // air pressure (hPa)
		20)        // air temperature (°C)

	fmt.Printf("%.3f %.3f\n", azimuth, zenithAngle)
	// Output: 162.233 29.990
}

func TestTimeConv(t *testing.T) {
	d := time.Date(2012, 1, 1, 12, 15, 45, 0, time.FixedZone("CET", 60*60))

	dt := calcT(d)

	if dt != -17531.53072879125 {
		t.Error("time conversion failed, value is ", dt)
	}
}

func TestSampleData(t *testing.T) {
	const tolerance = 0.001

	var cReferenceData = [][]float64{
		{1, 2.77132, -2.59916},
		{2, 2.63293, -2.12723},
		{3, 2.45419, -1.83734},
		{4, 2.26213, -1.63029},
		{5, 2.0676, -1.45852},
		{6, 1.87655, -1.29906},
		{7, 1.6938, -1.13844},
		{8, 1.52046, -0.96688},
		{9, 1.37407, -0.776354},
		{10, 1.25365, -0.561006},
		{11, 1.17059, -0.319657},
		{12, 1.13381, -0.0591845},
		{13, 1.14799, 0.205276},
		{14, 1.21126, 0.456612},
		{15, 1.3162, 0.683898},
		{16, 1.45273, 0.88506},
		{17, 1.61603, 1.06429},
		{18, 1.79341, 1.22882},
		{19, 1.98131, 1.38759},
		{20, 2.17486, 1.55202},
		{21, 2.36885, 1.73975},
		{22, 2.55533, 1.98455},
		{23, 2.71668, 2.35974},
		{24, 2.80795, 2.97419},
	}

	for _, row := range cReferenceData {
		d := time.Date(2012, 1, 1, int(row[0]), 0, 0, 0, time.FixedZone("CET", 60*60))

		refAzi := convertAzimuth(row[2])
		refZen := toDeg(row[1])

		azi, zen := Grena3(d, toDeg(0.73117), toDeg(0.21787), 65, 1000, 20)

		aziDiff := math.Abs(refAzi - azi)
		zenDiff := math.Abs(refZen - zen)

		if aziDiff > tolerance || zenDiff > tolerance {
			t.Error("azi/zen different. ref: ", refAzi, refZen, " vs. actual ", azi, zen)
		}

	}

}
