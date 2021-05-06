package abstractfactory

// CruiseMotorbike ...
type CruiseMotorbike struct{}

// NumWheels ...
func (c *CruiseMotorbike) NumWheels() int {
	return 2
}

// NumSeats ...
func (c *CruiseMotorbike) NumSeats() int {
	return 2
}

// GetMotorbikeType ...
func (c *CruiseMotorbike) GetMotorbikeType() int {
	return CruiseMotorbikeType
}
