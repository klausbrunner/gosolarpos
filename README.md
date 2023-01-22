gosolarpos
==========

![](https://github.com/KlausBrunner/gosolarpos/workflows/Go/badge.svg)

Package gosolarpos contains functions to find topocentric solar coordinates, i.e. the sun’s position on the sky at a given date, latitude, and longitude (and other parameters).

This is a partial Golang port of https://github.com/KlausBrunner/solarpositioning. Currently, only the [Grena/ENEA](http://dx.doi.org/10.1016/j.solener.2012.01.024) #3 algorithm is included - which should be fine for most applications, and preferable to many of the unattributed sun position calculation snippets floating around the Web.

Usage
-----

To use this module:
```Go
import (
    "github.com/KlausBrunner/gosolarpos"
)

azimuthDegrees, zenithDegrees := gosolarpos.Grena3(time.Now(),
		52.521667, // latitude (degrees)
		13.413333, // longitude (degrees)
		68,        // Delta T (s) - if unknown, use EstimateDeltaT()
		1000,      // air pressure (hPa)
		20)        // air temperature (°C)

println(azimuthDegrees, zenithDegrees)
```

A simple CLI is also available:

```console
$ cd cmd/sol
$ go install
$ sol --lat 48.266667 --lon -116.566667                 
 using current time: 2023-01-21 13:35:26.139919 -0800 PST m=+0.000353835
 estimating delta-T: 73.310637
 using standard sea-level pressure: 1013.250000 hPa
azumith: 204.254092   zenith: 71.457228
```
