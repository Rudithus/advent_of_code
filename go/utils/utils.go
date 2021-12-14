package utils

type Empty struct{}
type Problem interface {
	SolveQ1([]byte) int
	SolveQ2([]byte) int
	Path() string
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}
