package matrix_util

import (
	"fmt"
	"testing"
)

func TestProductMatrixVector(t *testing.T) {

	// [1,2,3]
	// [4,5,6]
	// [7,8,9]
	// [1,2,3]
	inputMatrix := [][]float64{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
		{1, 2, 3},
	}
	// [1]
	// [2]
	// [3]
	inputVector := []float64{1, 2, 3}

	got := ProductMatrixVector(inputMatrix, inputVector)

	expected := []float64{14, 32, 50, 14}

	for i, value := range expected {
		if got[i] != value {
			t.Fail()
			t.Logf(fmt.Sprintf("expected: %.2f - got %.2f ", expected[i], got[i]))
		}
	}

}
