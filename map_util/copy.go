package map_util

func CopyMap[k ~string, v any](m map[k]v) (copied map[k]v) {
	copied = make(map[k]v)
	for key, value := range m {
		copied[key] = value
	}
	return copied
}
