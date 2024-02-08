package vector_util

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMultiplyValues(t *testing.T) {

	vector1 := []float64{
		1, 2, 3,
	}
	vector2 := []float64{
		4, 5, 6,
	}

	got := MultiplyValuesFloats(vector1, vector2)
	expected := []float64{4, 10, 18}

	for i, value := range got {
		if value != expected[i] {
			t.Fail()
			t.Logf(fmt.Sprintf("expected: %.2f - got %.2f ", expected[i], got[i]))
		}
	}

}

func TestAddVectors(t *testing.T) {
	vector1 := []float64{
		1, 2, 3,
	}
	vector2 := []float64{
		4, 5, 6,
	}

	got := AddAToB(vector1, vector2)
	expected := []float64{5, 7, 9}

	for i, value := range got {
		if value != expected[i] {
			t.Fail()
			t.Logf(fmt.Sprintf("expected: %.2f - got %.2f ", expected[i], got[i]))
		}
	}

}

func TestSum(t *testing.T) {
	vector1 := []float64{
		1, 2, 3,
	}
	got := Sum(vector1)
	expected := 6.0
	if got != float64(expected) {
		t.Fail()
		t.Logf(fmt.Sprintf("expected: %.2f - got %.2f ", expected, got))
	}
}

func TestReluAndNormaliseBySum(t *testing.T) {

	vector1 := []float64{
		-1, 1, 3,
	}

	got := ReluAndNormaliseBySum(vector1)
	expected := []float64{0, 0.25, 0.75}
	for i, value := range got {
		if value != expected[i] {
			t.Fail()
			t.Logf(fmt.Sprintf("expected: %.2f - got %.2f ", expected[i], got[i]))
		}
	}
}

func TestDistanceBetweenPoints(t *testing.T) {

	given1 := []float64{1, 2}
	given2 := []float64{3, 4}
	want := 2.8284271247462

	gotPtr, err := DistanceBetweenPoints(given1, given2)
	assert.NoError(t, err)
	assert.InDelta(t, want, *gotPtr, 0.00005)

	given1 = []float64{5, 6, 2}
	given2 = []float64{-7, 11, -13}
	want = 19.849433

	gotPtr, err = DistanceBetweenPoints(given1, given2)
	assert.NoError(t, err)
	assert.InDelta(t, want, *gotPtr, 0.005)
}

func TestNormalise(t *testing.T) {

	in := []float64{1, 2, 3}
	want := []float64{
		0.2672612419124244,
		0.5345224838248488,
		0.8017837257372732,
	}

	got := Normalise(in)

	assert.InDelta(t, want[0], got[0], 0.000005)
	assert.InDelta(t, want[1], got[1], 0.000005)
	assert.InDelta(t, want[2], got[2], 0.000005)

	in = []float64{4, 5}
	want = []float64{
		0.6246950475544243,
		0.7808688094430304,
	}

	got = Normalise(in)

	assert.InDelta(t, want[0], got[0], 0.000005)
	assert.InDelta(t, want[1], got[1], 0.000005)
}

func TestVectorBetweenPoints(t *testing.T) {

	given1 := []float64{1, 2}
	given2 := []float64{3, 4}

	want := []float64{2, 2}

	got, err := VectorBetweenPoints(given1, given2)

	assert.NoError(t, err)
	assert.InDelta(t, want[0], got[0], 0.00005)
	assert.InDelta(t, want[1], got[1], 0.00005)

	given1 = []float64{0, 1}
	given2 = []float64{0, 0}

	want = []float64{0, -1}

	got, err = VectorBetweenPoints(given1, given2)

	assert.NoError(t, err)
	assert.InDelta(t, want[0], got[0], 0.00005)
	assert.InDelta(t, want[1], got[1], 0.00005)

}
