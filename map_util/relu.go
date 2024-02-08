package map_util

import "math"

func Relu(m map[string]float64) map[string]float64 {
	for k, v := range m {
		m[k] = math.Max(0, v)
	}
	return m
}

func Balance(m map[string]float64, max, min float64) map[string]float64 {
	for k, v := range m {
		if v > max {
			m[k] = max
		} else if v < min {
			m[k] = min
		}
	}
	return m
}
