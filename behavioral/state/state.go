package state

import "fmt"

// TrafficLightContext is the context class that can change its state like a traffic light.
type TrafficLightContext struct {
	State TrafficLightState
}

// SetState changes the current state of the traffic light.
func (t *TrafficLightContext) SetState(state TrafficLightState) {
	t.State = state
}

// Change requests the traffic light to change its current state.
func (t *TrafficLightContext) Change() {
	t.State.Change(t)
}

// TrafficLightState defines an interface for the states of the traffic light.
type TrafficLightState interface {
	Change(*TrafficLightContext)
}

// RedLightState is a concrete state representing the red light of a traffic signal.
type RedLightState struct{}

// Change method for RedLightState, changes the state to GreenLightState.
func (r *RedLightState) Change(t *TrafficLightContext) {
	t.SetState(&GreenLightState{})
	fmt.Println("Green light is on. Go!")
}

// GreenLightState is a concrete state representing the green light of a traffic signal.
type GreenLightState struct{}

// Change method for GreenLightState, changes the state to YellowLightState.
func (g *GreenLightState) Change(t *TrafficLightContext) {
	t.SetState(&YellowLightState{})
	fmt.Println("Yellow light is on. Slow down and prepare to stop!")
}

// YellowLightState is a concrete state representing the yellow light of a traffic signal.
type YellowLightState struct{}

// Change method for YellowLightState, changes the state to RedLightState.
func (y *YellowLightState) Change(t *TrafficLightContext) {
	t.SetState(&RedLightState{})
	fmt.Println("Red light is on. Get ready to go!")
}
