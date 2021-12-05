package q1

type Submarine struct {
	Pos Position
}

func (sub *Submarine) Down(value int) {
	sub.Pos.Y += value
}
func (sub *Submarine) Up(value int) {
	sub.Pos.Y -= value
}
func (sub *Submarine) Forward(value int) {
	sub.Pos.X += value
}

type Position struct {
	X int
	Y int
}
