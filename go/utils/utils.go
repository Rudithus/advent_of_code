package utils

type ReadLine func() (line string, eof bool)

type Problem interface {
	SolveQ1(ReadLine) string
	SolveQ2(ReadLine) string
	Input() string
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func Sum(arr []int) int {
	sum := 0
	for _, i := range arr {
		sum += i
	}
	return sum
}

func CompareInt(one, other int) int {
	if one > other {
		return 1
	}
	if one < other {
		return -1
	}

	return 0
}
