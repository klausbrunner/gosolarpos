package main

import (
	"fmt"
	"log"
	"math"
	"time"

	flag "github.com/spf13/pflag"

	"github.com/KlausBrunner/gosolarpos"
)

const (
	flagLat          = "lat"
	flagLon          = "lon"
	flagDate         = "date"
	flagDeltaT       = "delta-t"
	flagPressureHPa  = "pressure-hpa"
	flagPressureInHg = "pressure-inhg"
	flagTemperatureC = "temperature-c"
	flagTemperatureF = "temperature-f"
)

func main() {
	flag.Float64(flagLat, math.MaxFloat64,
		"(required) latitude")
	flag.Float64(flagLon, math.MaxFloat64,
		"(required) longitude")
	flag.String(flagDate, "",
		"(optional) date")
	flag.Float64(flagDeltaT, 0,
		"(optional) delta-T value")
	flag.Float64(flagPressureHPa, 1013.25,
		"(optional) barometric pressure as hectopascals (hPa)")
	flag.Float64(flagPressureInHg, 0,
		"(optional) barometric pressure as inches of mercury (inHg)")
	flag.Float64(flagTemperatureC, 15,
		"(optional) temperature as degrees Celsius (C)")
	flag.Float64(flagTemperatureF, 0,
		"(optional) temperature as degrees Fahrenheit (F)")
	flag.Parse()

	var lat float64
	if f := flag.Lookup(flagLat); f.Changed {
		v, err := flag.CommandLine.GetFloat64(f.Name)
		if err != nil {
			panic(err)
		}
		if 90 < lat || lat < -90 {
			log.Fatalf("latitude value %.6f invalid", lat)
		}
		lat = v
	} else {
		log.Fatalln("latitude value missing")
	}

	var lon float64
	if f := flag.Lookup(flagLon); f.Changed {
		v, err := flag.CommandLine.GetFloat64(f.Name)
		if err != nil {
			panic(err)
		}
		if 180 < lon || lon < -180 {
			log.Fatalf("longitude value %.6f invalid", lon)
		}
		lon = v
	} else {
		log.Fatalln("longitude value missing")
	}

	var date time.Time
	if f := flag.Lookup(flagDate); f.Changed {
		v, err := flag.CommandLine.GetString(f.Name)
		if err != nil {
			panic(err)
		}
		d, err := time.Parse(time.RFC3339, v)
		if err != nil {
			log.Fatalf("failed to parse date '%s':\n  %s", v, err.Error())
		}
		date = d
	} else {
		date = time.Now()
		fmt.Printf(" using current time: %s\n", date.String())
	}

	var deltaT float64
	if f := flag.Lookup(flagDeltaT); f.Changed {
		v, err := flag.CommandLine.GetFloat64(f.Name)
		if err != nil {
			panic(err)
		}
		deltaT = v
	} else {
		deltaT = gosolarpos.EstimateDeltaT(date)
		fmt.Printf(" estimating delta-T: %.6f\n", deltaT)
	}

	var pressureHPa float64
	if f := flag.Lookup(flagPressureInHg); f.Changed {
		v, err := flag.CommandLine.GetFloat64(f.Name)
		if err != nil {
			panic(err)
		}
		pressureHPa = v * 338500
	} else {
		f = flag.Lookup(flagPressureHPa)
		v, err := flag.CommandLine.GetFloat64(f.Name)
		if err != nil {
			panic(err)
		}
		pressureHPa = v
		if !f.Changed {
			fmt.Printf(" using standard sea-level pressure: %.6f hPa\n", pressureHPa)
		}
	}

	var temperatureC float64
	if f := flag.Lookup(flagTemperatureF); f.Changed {
		v, err := flag.CommandLine.GetFloat64(f.Name)
		if err != nil {
			panic(err)
		}
		temperatureC = (v - 32) * 5 / 9
	} else {
		f = flag.Lookup(flagTemperatureC)
		v, err := flag.CommandLine.GetFloat64(f.Name)
		if err != nil {
			panic(err)
		}
		temperatureC = v
		if f.Changed {
			fmt.Printf(" using standard sea-level temperature: %.6f C\n", temperatureC)
		}
	}

	azimuth, zenith := gosolarpos.Grena3(date, lat, lon, deltaT, pressureHPa, temperatureC)
	fmt.Printf("azumith: %.6f   zenith: %.6f\n", azimuth, zenith)
}
