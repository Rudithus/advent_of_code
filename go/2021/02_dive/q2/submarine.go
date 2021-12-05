package q2

type Submarine struct {
	Pos Position
	Aim int
}

func (sub *Submarine) Down(value int) {
	sub.Aim += value
}
func (sub *Submarine) Up(value int) {
	sub.Aim -= value
}
func (sub *Submarine) Forward(value int) {
	sub.Pos.X += value
	sub.Pos.Y += sub.Aim * value
}

type Position struct {
	X int
	Y int
}
