# Abstract Factory Pattern

The Abstract Factory design pattern is a new layer of grouping to achieve a bigger (and
more complex) composite object, which is used through its interfaces. The idea behind
grouping objects in families and grouping families is to have big factories that can be
interchangeable and can grow more easily. (Contreras, 2017)

## Implementation

The implementation uses vehicle factory as an example, where vehicle factory is _a super factory_ or _factory of factories_.

### Types

Super factory is the following:

```go
package abstractfactory

import (
	"fmt"
)

// VehicleFactory is a super factory, which returns another factory
type VehicleFactory interface {
	NewVehicle(v int) (Vehicle, error)
}

```

The following _Vehicle_ factory is a factory for _Family Car_, _Luxury Car_, _Sport Motorbike_ and _Cruise Motorbike_ that are **`families of related objects`**.
```go
// Vehicle ...
type Vehicle interface {
	NumWheels() int
	NumSeats() int
}
```

## Usage

This pattern is commonly used in many applications and libraries, such as cross-platform GUI libraries. Think of a button, a generic object, and button factory that provides you with a factory for Microsoft Windows buttons while you have another
factory for Mac OS X buttons. You don't want to deal with the implementation details of
each platform, but you just want to implement the actions for some specific behavior raised by a button. (Contreras, 2017)