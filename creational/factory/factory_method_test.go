package factory

import (
	"testing"
)

func TestCreateVehicle(t *testing.T) {
	tests := []struct {
		vehicleType VehicleType
		expected    string
		shouldError bool
	}{
		{CarType, "Driving a Car", false},
		{BikeType, "Riding a Bike", false},
		{VehicleType(999), "", true},
	}

	for _, test := range tests {
		vehicle, err := CreateVehicle(test.vehicleType)
		if test.shouldError {
			if err == nil {
				t.Errorf("expected an error for vehicle type %v, but got none", test.vehicleType)
			}
		} else {
			if err != nil {
				t.Errorf("did not expect an error for vehicle type %v, but got %v", test.vehicleType, err)
			}
			if vehicle.Drive() != test.expected {
				t.Errorf("expected vehicle drive to be %v, but got %v", test.expected, vehicle.Drive())
			}
		}
	}
}
