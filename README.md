gosolarpos
==========

![](https://github.com/klausbrunner/gosolarpos/workflows/Go/badge.svg)

Package gosolarpos contains functions to find topocentric solar coordinates, i.e. the sun’s position on the sky at a given date, latitude, and longitude (and other parameters).

This is a partial Golang port of https://github.com/klausbrunner/solarpositioning. Currently, only the [Grena/ENEA](http://dx.doi.org/10.1016/j.solener.2012.01.024) #3 algorithm is included - which should be fine for most applications, and preferable to many of the unattributed sun position calculation snippets floating around the Web.

Usage
-----

To use this module:
```Go
import (
    "github.com/klausbrunner/gosolarpos"
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
 using current time: 2023-01-22 09:46:27.86356 +0100 CET m=+0.000386917
 estimating delta-T: 70.81
 using standard sea-level pressure: 1013.25 hPa
azimuth: 23.252482°   zenith: 149.786614°
```
