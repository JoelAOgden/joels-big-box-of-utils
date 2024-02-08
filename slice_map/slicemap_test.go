package slice_map

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSliceMapPointers(t *testing.T) {

	type testSliceObject struct {
		I int
	}

	test := New[testSliceObject]()

	test.Put("1", testSliceObject{1})

	get := test.Get("1")
	fmt.Printf("%p\n", get)
	(*get).I = 2

	get2 := test.Get("1")
	fmt.Printf("%p\n", get2)
	get3 := test.Get("1")
	fmt.Printf("%p\n", get3)
	assert.Equal(t, 2, (*test.Get("1")).I)
}
