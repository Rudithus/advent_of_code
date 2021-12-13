package utils

func CompareInt(one, other int) int {
	if one > other {
		return 1
	}
	if one < other {
		return -1
	}

	return 0
}
func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
