package composite

// Athlete ...
type Athlete struct{}

// Train ...
func (a *Athlete) Train() string {
	return "training"
}

// CompositeSwimmerA - first method with direct composition
type CompositeSwimmerA struct {
	MyAthlete Athlete
	MySwim    func() string
}

// Swim ...
func Swim() string {
	return "swimming"
}

// Animal ...
type Animal struct{}

// Eat ...
func (r *Animal) Eat() string {
	return "eating"
}

// Shark - second method with embedded composition
type Shark struct {
	Animal
	Swim func() string
}

// Swimmer ...
type Swimmer interface {
	Swim() string
}

// Trainer ...
type Trainer interface {
	Train() string
}

// SwimmerImpl ...
type SwimmerImpl struct{}

// Swim ...
func (s *SwimmerImpl) Swim() string {
	return "swimming"
}

// CompositeSwimmerB - third method with SwimmerImpl type
type CompositeSwimmerB struct {
	Trainer
	Swimmer
}
