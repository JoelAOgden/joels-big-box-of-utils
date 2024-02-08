package vector_util

import (
	"errors"
	"fmt"
	"github.com/JoelAOgden/joels-big-box-of-utils/types"
	"math"
	"math/rand"
)

func FloatsToInts(input []float64) []int {
	output := make([]int, len(input))
	for i, value := range input {
		output[i] = int(value)
	}
	return output
}

func FillVectorWithRandom(input *[]float64, min float64, max float64) {
	for i := range *input {
		(*input)[i] = (rand.Float64() * (max - min)) + min
	}
}

func MultiplyByScale[k types.Number](in []k, scale float64) []float64 {
	out := make([]float64, len(in))
	for i, value := range in {
		out[i] = float64(value) * scale
	}
	return out
}

// Multiplies two vector values such that n+1 = xn * yn
//
// E.g [1,2,3] * [9,8,7] = [1*9, 2*8, 3*7]
//
// See: Hadamard Product
func MultiplyValuesFloats(vector1 []float64, vector2 []float64) []float64 {

	if len(vector1) != len(vector2) {
		panic(errors.New("arrays not the same length"))
	}

	output := make([]float64, len(vector1))
	for i := range vector1 {
		output[i] = vector1[i] * vector2[i]
	}

	return output
}

// Multiplies two vector values such that n+1 = xn * yn
//
// E.g [1,2,3] * [9,8,7] = [1*9, 2*8, 3*7]
//
// See: Hadamard Product
func MultiplyValuesInts(vector1 []int, vector2 []int) []int {

	if len(vector1) != len(vector2) {
		panic(errors.New("arrays not the same length"))
	}

	output := make([]int, len(vector1))
	for i := range vector1 {
		output[i] = vector1[i] * vector2[i]
	}

	return output
}

// Multiplies two vector values such that n+1 = xn * yn
//
// E.g [1,2,3] * [9,8,7] = [1*9, 2*8, 3*7]
//
// See: Hadamard Product
func MultiplyValuesIntsAndFloats(vector1 []int, vector2 []float64) []float64 {

	if len(vector1) != len(vector2) {
		panic(errors.New("arrays not the same length"))
	}

	output := make([]float64, len(vector1))
	for i := range vector1 {
		output[i] = float64(vector1[i]) * vector2[i]
	}

	return output
}

// vector addition
//
// e.g. [1,2] + [3,4] = [1+3, 2+4] = [4,6]
func AddAToB(A []float64, B []float64) []float64 {
	if len(A) != len(B) {
		panic(errors.New("slices not the same length"))
	}

	output := make([]float64, len(A))
	for i := range A {
		output[i] = A[i] + B[i]
	}

	return output
}

// Sum the vector
//
// E.g [1,2] = 1 + 2
func Sum(input []float64) float64 {
	sum := 0.0
	for _, value := range input {
		sum += value
	}
	return sum
}

// Relu the vector and normalise by the sum
//
// Relu(x) = max(0,x)
//
// Normalised value = x / sum
//
// E.g. [-1,2] = [0,1]
//
// See: Rectifier (neural networks)
func ReluAndNormaliseBySum(input []float64) []float64 {

	input = Relu(input)

	return NormaliseBySum(input)
}

func Relu(input []float64) []float64 {
	for i, value := range input {
		input[i] = math.Max(value, 0)
	}

	return input
}

func Normalise(input []float64) []float64 {

	zeroes := Zeroes(len(input))
	distancePtr, err := DistanceBetweenPoints(zeroes, input)
	if err != nil {
		panic("explosion lol") // todo: panics bad
	}

	distance := *distancePtr
	out := MultiplyByScale(input, 1./distance)

	return out
}

// [0,0,0,0...n]
func Zeroes(n int) []float64 {

	zeroVector := make([]float64, n)
	for i := 0; i < n; i++ {
		zeroVector[i] = 0.0
	}

	return zeroVector

}

// Normalised based on sum not unit vector
// ie. [1,3] = [0.25, 0.75]
// NOT [1,3] = [0.32, 0.94]
func NormaliseBySum(input []float64) []float64 {

	sum := 0.0
	for _, value := range input {
		sum += value
	}

	output := make([]float64, len(input))
	for i, value := range input {
		if sum <= 0 {
			output[i] = 1.0 / float64(len(input))
		} else {
			output[i] = value / sum
		}
	}

	return output
}

// vector subtraction
//
// e.g. [1,2] + [3,4] = [1-3, 2-4] = [-2,-2]
func SubtractBFromA(a []float64, b []float64) ([]float64, error) {
	if len(a) != len(b) {
		return nil, fmt.Errorf("slices not the same length")
	}

	out := make([]float64, len(a))
	for i := range out {
		out[i] = a[i] - b[i]
	}

	return out, nil
}

// finds vector vector between two points
//
// vector is not normalised
func VectorBetweenPoints(a []float64, b []float64) ([]float64, error) {
	if len(a) != len(b) {
		return nil, fmt.Errorf("slices not the same length")
	}

	vector, err := SubtractBFromA(b, a)
	if err != nil {
		return nil, err
	}

	return vector, nil

}

// finds direction vector between two points
//
// direction is normalised
func DirectionVectorBetweenPoints(a []float64, b []float64) ([]float64, error) {
	if len(a) != len(b) {
		return nil, fmt.Errorf("slices not the same length")
	}

	vector, err := VectorBetweenPoints(a, b)
	if err != nil {
		return nil, err
	}
	return Normalise(vector), nil

}

// finds the distance between two points
//
// n dimensional
func DistanceBetweenPoints(a []float64, b []float64) (*float64, error) {
	if len(a) != len(b) {
		return nil, fmt.Errorf("slices not the same length")
	}
	sum := 0.
	for i := range a {
		v := b[i] - a[i]
		sum += math.Pow(v, 2)
	}

	distance := math.Sqrt(sum)
	return &distance, nil
}
