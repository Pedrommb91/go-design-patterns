package factory

import "fmt"

// Vehicle is the interface that all concrete vehicles should implement
type Vehicle interface {
	Drive() string
}

// Car is a concrete implementation of the Vehicle interface
type Car struct{}

func (c *Car) Drive() string {
	return "Driving a Car"
}

// Bike is a concrete implementation of the Vehicle interface
type Bike struct{}

func (b *Bike) Drive() string {
	return "Riding a Bike"
}

// VehicleType is an enumeration of the different types of vehicles
type VehicleType int

const (
	CarType VehicleType = iota
	BikeType
)

// CreateVehicle is the factory method that creates vehicles based on the given type
func CreateVehicle(vehicleType VehicleType) (Vehicle, error) {
	switch vehicleType {
	case CarType:
		return &Car{}, nil
	case BikeType:
		return &Bike{}, nil
	default:
		return nil, fmt.Errorf("unknown vehicle type")
	}
}
