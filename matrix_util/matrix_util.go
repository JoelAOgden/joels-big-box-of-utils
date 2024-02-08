package matrix_util

import (
	"fmt"
	"github.com/JoelAOgden/joels-big-box-of-utils/vector_util"
)

func ProductMatrixVector(matrix [][]float64, array []float64) []float64 {
	if len(matrix[0]) != len(array) {
		panic(fmt.Errorf("incorrect matrix and vector lengths: %v, %v", len(matrix[0]), len(array)))
	}

	output := make([]float64, len(matrix))

	for i, firstArray := range matrix {
		multipliedArray := vector_util.MultiplyValuesFloats(firstArray, array)
		summedArray := vector_util.Sum(multipliedArray)
		output[i] = summedArray
	}

	return output
}
