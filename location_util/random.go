package location_util

import "math/rand"

func RandomLocation(min float64, max float64) Location {
	return Location{
		X: int((rand.Float64() * (min + max)) - min),
		Y: int((rand.Float64() * (min + max)) - min),
	}
}
