package vector_util

import "math/rand"

func RandomLocation(min float64, max float64) Location {
	return Location{
		X: int((rand.Float64() * (min + max)) - min),
		Y: int((rand.Float64() * (min + max)) - min),
	}
}

type Location struct {
	X, Y int
}

func (l *Location) SetLocation(x, y int) {
	l.X = x
	l.Y = y
}

func (l *Location) ToVector() [2]float64 {
	return [2]float64{
		float64(l.X), float64(l.Y),
	}
}

type DesiredLocation struct {
	Location Location
}

func (d DesiredLocation) GetLocation() Location {
	return d.Location
}
