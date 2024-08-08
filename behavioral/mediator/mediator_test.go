package mediator

import (
	"testing"
)

func TestControlTower_CanLand(t *testing.T) {
	controlTower := NewControlTower()
	plane1 := NewPassengerPlane(controlTower, 1)
	plane2 := NewPassengerPlane(controlTower, 2)

	if !controlTower.CanLand(plane1) {
		t.Error("Expected plane to be able to land when the runway is free")
	}

	if controlTower.CanLand(plane2) {
		t.Error("Expected plane not to be able to land when the runway is not free")
	}
}

func TestControlTower_NotifyAboutTakeoff(t *testing.T) {
	controlTower := NewControlTower()
	plane1 := NewPassengerPlane(controlTower, 1)
	plane2 := NewPassengerPlane(controlTower, 2)

	if !controlTower.CanLand(plane1) {
		t.Error("Expected plane1 to be able to land when the runway is free")
	}
	if controlTower.CanLand(plane2) {
		t.Error("Expected plane2 not to be able to land when the runway is not free")
	}
	plane1.Land()
	controlTower.NotifyAboutTakeoff()

	if controlTower.isRunwayFree {
		t.Error("Expected runway to be occupied after notifying about takeoff when there is a plane in the queue")
	}
}

func TestPassengerPlane_Land(t *testing.T) {
	controlTower := NewControlTower()
	plane := NewPassengerPlane(controlTower, 1)

	if !plane.Land() {
		t.Error("Expected plane to be able to land when the runway is free")
	}
}

func TestPassengerPlane_Takeoff(t *testing.T) {
	controlTower := NewControlTower()
	plane := NewPassengerPlane(controlTower, 1)

	controlTower.CanLand(plane)
	plane.Land()
	plane.Takeoff()

	if !controlTower.isRunwayFree {
		t.Error("Expected runway to be free after plane takeoff")
	}
}
