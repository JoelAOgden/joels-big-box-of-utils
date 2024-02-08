package slice_util

func RemoveIndex[T any](s []T, index int) []T {
	ret := make([]T, 0)
	ret = append(ret, s[:index]...)
	if index >= len(s) {
		return ret
	}

	final := append(ret, s[index+1:]...)
	return final
}
