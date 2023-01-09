gosolarpos
==========

![](https://github.com/KlausBrunner/gosolarpos/workflows/Go/badge.svg)

Package gosolarpos contains functions to find topocentric solar coordinates, i.e. the sun’s position on the sky at a given date, latitude, and longitude (and other parameters).

This is a partial Golang port of https://github.com/KlausBrunner/solarpositioning. Currently, only the [Grena/ENEA](http://dx.doi.org/10.1016/j.solener.2012.01.024) #3 algorithm is included - which should be fine for most applications, and preferable to many of the unattributed sun position calculation snippets floating around the Web.

Usage
-----

```Go
import (
    "github.com/KlausBrunner/gosolarpos"
)

azimuth, zenithAngle := Grena3(time.Now(),
		52.521667, // latitude (degrees)
		13.413333, // longitude (degrees)
		68,        // Delta T (s) - if unknown, use EstimateDeltaT()
		1000,      // air pressure (hPa)
		20)        // air temperature (°C)
```
