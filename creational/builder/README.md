# Builder Pattern

The Builder pattern helps us construct complex objects without directly instantiating their struct, or writing the logic they require. Imagine an object that could have dozens of fields that are more complex structs themselves. Now imagine that you have many objects with these characteristics, and you could have more. We don't want to write the logic to create all these objects in the package that just needs to use the objects. (Contreras, 2017)

## Implementation

Creating a vehicle (the product) could be considered as an example here. The process (widely described as the algorithm) of creating a vehicle (the product) is more or less the same for every kind of vehicle–choose vehicle type, assemble the structure, place the wheels, and place the seats. If you think about it, you could build a car and a motorbike (two Builders) with this description, so we are reusing the description to create cars in manufacturing. (Contreras, 2017)

Builder interface for vehicle has methods _SetWheels_, _SetSeats_, _SetStructure_ that returns the Builder itself. At the end, the method Build to finish the build process.

```go
type Builder interface {
	SetWheels() Builder
	SetSeats() Builder
	SetStructure() Builder
	Build() VehicleProduct
}

```

_Builder_ is implemented by _CarBuilder_, _BusBuilder_, _BikeBuilder_, each having different implementations of interface's methods. For example, here is the _CarBuilder_:

```go
type CarBuilder struct {
	v VehicleProduct
}

func (c *CarBuilder) SetWheels() Builder {
	c.v.Wheels = 4
	return c
}

func (c *CarBuilder) SetSeats() Builder {
	c.v.Seats = 5
	return c
}

func (c *CarBuilder) SetStructure() Builder {
	c.v.Structure = "Car"
	return c
}

func (c *CarBuilder) Build() VehicleProduct {
	return c.v
}
```

## Final discussions

The Builder design pattern helps us maintain an unpredictable number of products by using a common construction algorithm that is used by the director. The construction process is always abstracted from the user of the product. 

At the same time, having a defined construction pattern helps when a newcomer to our source code needs to add a new product to the pipeline. The BuildProcess interface specifies what he must comply to be part of the possible builders.

However, try to avoid the Builder pattern when you are not completely sure that the
algorithm is going to be more or less stable because any small change in this interface will affect all your builders and it could be awkward if you add a new method that some of your builders need and others Builders do not. (Contreras, 2017)
