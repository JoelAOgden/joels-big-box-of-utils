package location_util

import (
	"fmt"
	"github.com/JoelAOgden/joels-big-box-of-utils/vector_util"
)

// Location is 2d point vector in int format (for easier drawing)
type Location struct {
	X, Y int
}

func (l *Location) SetLocation(x, y int) {
	l.X = x
	l.Y = y
}

func (l *Location) ToVector() []float64 {
	return []float64{
		float64(l.X), float64(l.Y),
	}
}

func (l *Location) LocationTowardsLocation(distance int, location Location) (*Location, error) {

	dir, err := vector_util.DirectionVectorBetweenPoints(
		l.ToVector(),
		location.ToVector())
	if err != nil {
		return nil, err
	}

	if len(dir) != 2 {
		return nil, fmt.Errorf("incorrect direction vector length: %v", len(dir))
	}

	xOffset := int(dir[0] * float64(distance))
	yOffset := int(dir[1] * float64(distance))

	return &Location{
		X: l.X + xOffset,
		Y: l.Y + yOffset,
	}, nil

}
