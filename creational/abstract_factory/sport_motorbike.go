package abstractfactory

// SportMotorbike ...
type SportMotorbike struct{}

// NumWheels ...
func (s *SportMotorbike) NumWheels() int {
	return 2
}

// NumSeats ...
func (s *SportMotorbike) NumSeats() int {
	return 1
}

// GetMotorbikeType ...
func (s *SportMotorbike) GetMotorbikeType() int {
	return SportMotorbikeType
}
