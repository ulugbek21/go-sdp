package abstractfactory

import (
	"fmt"
)

// VehicleFactory ...
type VehicleFactory interface {
	NewVehicle(v int) (Vehicle, error)
}

const (
	// CarFactoryType ...
	CarFactoryType       = 1
	// MotorbikeFactoryType ...
	MotorbikeFactoryType = 2
)

// BuildFactory ...
func BuildFactory(f int) (VehicleFactory, error) {
	switch f {
	case CarFactoryType:
		return new(CarFactory), nil
	case MotorbikeFactoryType:
		return new(MotorbikeFactory), nil
	default:
		return nil, fmt.Errorf("Factory with id %d not recognized", f)
	}
}

const (
	// LuxuryCarType ...
	LuxuryCarType = 1
	// FamilyCarType ...
	FamilyCarType = 2
)

// CarFactory ...
type CarFactory struct{}

// NewVehicle ...
func (c *CarFactory) NewVehicle(v int) (Vehicle, error) {
	switch v {
	case LuxuryCarType:
		return new(LuxuryCar), nil
	case FamilyCarType:
		return new(FamilyCar), nil
	default:
		return nil, fmt.Errorf("vehicle of type %d not recognized", v)
	}
}

const (
	// SportMotorbikeType ...
	SportMotorbikeType = 1
	// CruiseMotorbikeType ...
	CruiseMotorbikeType = 2
)

// MotorbikeFactory ...
type MotorbikeFactory struct{}

// NewVehicle ...
func (m *MotorbikeFactory) NewVehicle(v int) (Vehicle, error) {
	switch v {
	case SportMotorbikeType:
		return new(SportMotorbike), nil
	case CruiseMotorbikeType:
		return new(CruiseMotorbike), nil
	default:
		return nil, fmt.Errorf("vehicle of type %d not recognized", v)
	}
}
