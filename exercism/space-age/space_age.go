package space

import (
	"log"
	"math"
)

// Planet defines the type for all planets
type Planet string

// Age returns an earth year number to represent how old someone is on the given planet
func Age(seconds float64, planet Planet) float64 {
	planets := map[Planet]float64{
		"Mercury": 0.2408467,
		"Venus":   0.61519726,
		"Mars":    1.8808158,
		"Earth":   1.0,
		"Jupiter": 11.862615,
		"Saturn":  29.447498,
		"Uranus":  84.016846,
		"Neptune": 164.79132,
	}
	p, ok := planets[planet]
	if !ok {
		log.Fatal(planet, " is not part of the planets map.")
	}
	years := seconds / 31557600.0 / p
	return math.Round(years*100) / 100
}
