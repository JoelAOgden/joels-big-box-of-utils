package map_util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNormalise(t *testing.T) {
	in := map[string]float64{
		"a": 1,
		"b": 1,
	}

	want := map[string]float64{
		"a": 0.5,
		"b": 0.5,
	}

	got := Normalise(in)

	assert.Equal(t, want, got)

	in = map[string]float64{
		"a": 1000,
		"b": 2000,
		"c": 1000,
	}

	want = map[string]float64{
		"a": 0.25,
		"b": 0.5,
		"c": 0.25,
	}

	got = Normalise(in)

	assert.Equal(t, want, got)
}
