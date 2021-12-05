package q1

import (
	"testing"
)

func createsub() Submarine { return Submarine{Position{5, 5}} }

func TestUp(t *testing.T) {
	sub := createsub()
	expectedposition := Position{5, 0}
	sub.Up(5)

	if sub.Pos != expectedposition {
		t.Fail()
	}
}

func TestDown(t *testing.T) {
	sub := createsub()
	expectedposition := Position{5, 10}
	sub.Down(5)

	if sub.Pos != expectedposition {
		t.Fail()
	}
}

func TestForward(t *testing.T) {
	sub := createsub()
	expectedposition := Position{10, 5}
	sub.Forward(5)

	if sub.Pos != expectedposition {
		t.Fail()
	}
}
