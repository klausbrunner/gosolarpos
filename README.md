gosolarpos
==========

Package gosolarpos contains functions to find topocentric solar coordinates, i.e. the sun’s position on the sky at a given date, latitude, and longitude (and other parameters).

This is a partial, early-stage Golang port of https://github.com/KlausBrunner/solarpositioning. Currently, only the Grena3 algorithm is included (which is still useful for many applications, and preferrable to many of the unattributed sun position calculation snippets floating around the Web).

Usage
-----

```Go
azimuth, zenithAngle := Grena3(time.Now(),
		52.521667, // latitude (degrees)
		13.413333, // longitude (degrees)
		65,        // Delta T (s) - if unknown, use EstimateDeltaT()
		1000,      // air pressure (hPa)
		20)        // air temperature (°C)
```
