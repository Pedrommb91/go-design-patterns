package mediator

import "fmt"

// Mediator interface defines the communication protocol between different components.
type Mediator interface {
	CanLand(Aircraft) bool
	NotifyAboutTakeoff()
}

// ControlTower is a concrete mediator that coordinates how aircraft land and take off at an airport.
type ControlTower struct {
	isRunwayFree  bool
	aircraftQueue []Aircraft
}

// NewControlTower creates a new instance of ControlTower.
func NewControlTower() *ControlTower {
	return &ControlTower{
		isRunwayFree: true,
	}
}

// CanLand checks if the aircraft can land at the airport.
func (c *ControlTower) CanLand(aircraft Aircraft) bool {
	if c.isRunwayFree && len(c.aircraftQueue) == 0 {
		c.isRunwayFree = false
		return true
	}
	c.aircraftQueue = append(c.aircraftQueue, aircraft)
	return false
}

// NotifyAboutTakeoff notifies that an aircraft has taken off, so the runway is free.
func (c *ControlTower) NotifyAboutTakeoff() {
	if !c.isRunwayFree && len(c.aircraftQueue) > 0 {
		firstAircraft := c.aircraftQueue[0]
		c.aircraftQueue = c.aircraftQueue[1:]
		firstAircraft.PermitLanding()
	}
	if len(c.aircraftQueue) == 0 {
		c.isRunwayFree = true
	} else {
		c.isRunwayFree = false
	}
}

// Aircraft interface defines the methods for an aircraft interacting with the mediator.
type Aircraft interface {
	Land() bool
	Takeoff()
	PermitLanding()
}

// PassengerPlane is a concrete component that interacts with the mediator.
type PassengerPlane struct {
	mediator Mediator
	number   int
}

// NewPassengerPlane creates a new instance of PassengerPlane.
func NewPassengerPlane(mediator Mediator, number int) *PassengerPlane {
	return &PassengerPlane{
		mediator: mediator,
		number:   number,
	}
}

// Land tries to land at the airport and returns whether it's successful.
func (p *PassengerPlane) Land() bool {
	if p.mediator.CanLand(p) {
		fmt.Printf("Passenger Plane %d has landed.\n", p.number)
		return true
	} else {
		fmt.Printf("Passenger Plane %d is waiting to land.\n", p.number)
		return false
	}
}

// Takeoff notifies the mediator that the plane is taking off.
func (p *PassengerPlane) Takeoff() {
	fmt.Printf("Passenger Plane %d is taking off.\n", p.number)
	p.mediator.NotifyAboutTakeoff()
}

// PermitLanding is called by the mediator when the plane is allowed to land.
func (p *PassengerPlane) PermitLanding() {
	// Plane logic for being permitted to land
	fmt.Printf("Passenger Plane %d has permission to land.\n", p.number)
}
