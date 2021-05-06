package abstractfactory

// LuxuryCar ...
type LuxuryCar struct{}

// NumDoors ...
func (*LuxuryCar) NumDoors() int {
	return 4
}

// NumWheels ...
func (*LuxuryCar) NumWheels() int {
	return 4
}

// NumSeats ...
func (*LuxuryCar) NumSeats() int {
	return 5
}

