package gosolarpos

import (
	"math"
	"testing"
	"time"
)

func TestHistoricData(t *testing.T) {

	var referenceData = [][]float64{
		{100, 9600},
		{200, 8640},
		{300, 7680},
		{400, 6700},
		{500, 5710},
		{600, 4740},
		{700, 3810},
		{800, 2960},
		{900, 2200},
		{1000, 1570},
		{1100, 1090},
		{1200, 740},
		{1300, 490},
		{1400, 320},
		{1500, 200},
		{1600, 120},
		{1700, 9},
		{1750, 13},
		{1800, 14},
		{1850, 7},
		{1900, -3},
		{1950, 29},
		{1955, 31.1},
		{1990, 56.9},
		{2005, 64.7},
		{2015, 68.0},
	}

	for _, row := range referenceData {
		year := row[0]
		cal := time.Date(int(year), 1, 1, 0, 0, 0, 0, time.UTC)

		ref := row[1]
		est := EstimateDeltaT(cal)
		estDiffRatio := math.Abs(est-ref) / ref

		if estDiffRatio > 0.03 {
			t.Error("ref: ", ref, " actual: ", est)
		}

	}

}
