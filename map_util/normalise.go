package map_util

func Normalise(m map[string]float64) map[string]float64 {

	total := 0.
	for _, v := range m {
		total += v
	}

	// normalise
	for k, v := range m {
		m[k] = v / total
	}

	return m
}
