package uint_util

func SubtractUint(in uint, subtract uint) uint {

	if int(in)-int(subtract) <= 0 {
		return uint(0)
	}

	out := in - subtract

	return out

}
