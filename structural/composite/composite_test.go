package composite

import "testing"

func TestComposite(t *testing.T) {
	// First method with direct composition
	localSwim := Swim

	swimmer1 := CompositeSwimmerA{
		MySwim: localSwim,
	}

	t.Log("Swimmer 1", swimmer1.MyAthlete.Train())
	t.Log("Swimmer 1", swimmer1.MySwim())

	// Second Method with embedded composition
	shark := Shark{
		Swim: Swim,
	}

	t.Log("Shark", shark.Eat())
	t.Log("Shark", shark.Swim())

	swimmer2 := CompositeSwimmerB{
		&Athlete{},
		&SwimmerImpl{},
	}

	t.Log("Swimmer 2", swimmer2.Train())
	t.Log("Swimmer 2", swimmer2.Swim())
}
