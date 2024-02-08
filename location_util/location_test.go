package location_util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLocation_LocationTowardsLocation(t *testing.T) {

	location1 := Location{
		X: 1,
		Y: 1,
	}
	location2 := Location{
		X: 10,
		Y: 10,
	}

	want := &Location{
		X: 2,
		Y: 2,
	}
	got, err := location1.LocationTowardsLocation(2, location2)
	assert.NoError(t, err)
	assert.Equal(t, want, got)

	want = &Location{
		X: 3,
		Y: 3,
	}
	got, err = location1.LocationTowardsLocation(3, location2)
	assert.NoError(t, err)

	assert.Equal(t, want, got)

}
