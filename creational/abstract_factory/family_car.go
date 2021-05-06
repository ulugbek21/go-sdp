package abstractfactory

// FamilyCar ...
type FamilyCar struct{}

// NumDoors ...
func (*FamilyCar) NumDoors() int {
	return 5
}

// NumWheels ...
func (*FamilyCar) NumWheels() int {
	return 4
}

// NumSeats ...
func (*FamilyCar) NumSeats() int {
	return 5
}
