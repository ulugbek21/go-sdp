package builder

// Builder includes methods of Builder design pattern
type Builder interface {
	SetWheels() Builder
	SetSeats() Builder
	SetStructure() Builder
	Build() VehicleProduct
}

// VehicleProduct is the final object that we want to retrieve
type VehicleProduct struct {
	Wheels    int
	Seats     int
	Structure string
}

// ManufacturingDirector director variable is the one in charge of accepting the builders
// It has a Construct method that will use the builder that is stored in Manufacturing,
// and will reproduce the required steps.
// The SetBuilder method will allow us to change the builder that is being used in the Manufacturing director
type ManufacturingDirector struct {
	builder Builder
}

// Construct method will use builder that is stored in ManufacturingDirector
func (f *ManufacturingDirector) Construct() {
	f.builder.SetSeats().SetStructure().SetWheels()
}

// SetBuilder allows to change the builder that is being used in ManufacturingDirector
func (f *ManufacturingDirector) SetBuilder(b Builder) {
	f.builder = b
}

// CarBuilder implements Builder interface
type CarBuilder struct {
	v VehicleProduct
}

// SetWheels implements Builder's SetWheels() method
func (c *CarBuilder) SetWheels() Builder {
	c.v.Wheels = 4
	return c
}

// SetSeats implements Builder's SetSeats() method
func (c *CarBuilder) SetSeats() Builder {
	c.v.Seats = 5
	return c
}

// SetStructure implements Builder's SetStructure() method
func (c *CarBuilder) SetStructure() Builder {
	c.v.Structure = "Car"
	return c
}

// Build builds CarBuilder
func (c *CarBuilder) Build() VehicleProduct {
	return c.v
}

// BikeBuilder implements Builder interface
type BikeBuilder struct {
	v VehicleProduct
}

// SetWheels implements Builder's SetWheels() method
func (b *BikeBuilder) SetWheels() Builder {
	b.v.Wheels = 2
	return b
}

// SetSeats implements Builder's SetSeats() method
func (b *BikeBuilder) SetSeats() Builder {
	b.v.Seats = 2
	return b
}

// SetStructure implements Builder's SetStructure() method
func (b *BikeBuilder) SetStructure() Builder {
	b.v.Structure = "Motorbike"
	return b
}

// Build builds BikeBuilder
func (b *BikeBuilder) Build() VehicleProduct {
	return b.v
}

// BusBuilder implements Builder interface
type BusBuilder struct {
	v VehicleProduct
}

// SetWheels implements Builder's SetWheels() method
func (b *BusBuilder) SetWheels() Builder {
	b.v.Wheels = 4 * 2
	return b
}

// SetSeats implements Builder's SetSeats() method
func (b *BusBuilder) SetSeats() Builder {
	b.v.Seats = 30
	return b
}

// SetStructure implements Builder's SetStructure() method
func (b *BusBuilder) SetStructure() Builder {
	b.v.Structure = "Bus"
	return b
}

// Build builds BusBuilder
func (b *BusBuilder) Build() VehicleProduct {
	return b.v
}
