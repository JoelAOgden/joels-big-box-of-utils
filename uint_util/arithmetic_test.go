package uint_util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSubtractUint(t *testing.T) {

	in := uint(10)
	subtract := uint(5)

	got := SubtractUint(in, subtract)

	assert.Equal(t, uint(5), got)

	in = uint(10)
	subtract = uint(15)

	got = SubtractUint(in, subtract)

	assert.Equal(t, uint(0), got)

}
