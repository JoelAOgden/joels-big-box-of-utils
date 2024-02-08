package map_util

// Concatenate - concatenates two maps, map2 will overlap map_info 1
func Concatenate[k ~string, v any](map1 map[k]v, map2 map[k]v) (concatenated map[k]v) {
	concatenated = make(map[k]v)

	for k, v := range map1 {
		concatenated[k] = v
	}

	for k, v := range map2 {
		concatenated[k] = v
	}

	return concatenated
}
