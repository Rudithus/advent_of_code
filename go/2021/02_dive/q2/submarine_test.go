package q2

import (
	"testing"
)

func createsub() Submarine { return Submarine{Pos: Position{X: 5, Y: 5}, Aim: 5} }

func TestUp(t *testing.T) {
	sub := createsub()
	expectedAim := 0
	sub.Up(5)

	if sub.Aim != expectedAim {
		t.Fail()
	}
}

func TestDown(t *testing.T) {
	sub := createsub()
	expectedAim := 10
	sub.Down(5)

	if sub.Aim != expectedAim {
		t.Fail()
	}
}

func TestForward(t *testing.T) {
	sub := createsub()
	expectedposition := Position{10, 30}
	sub.Forward(5)

	if sub.Pos != expectedposition {
		t.Fail()
	}
}
