package state

import (
	"testing"
)

func TestTrafficLightStateChange(t *testing.T) {
	trafficLight := &TrafficLightContext{State: &RedLightState{}}

	// Red to Green
	trafficLight.Change()
	if _, ok := trafficLight.State.(*GreenLightState); !ok {
		t.Error("Expected traffic light state to be GreenLightState after change from RedLightState")
	}

	// Green to Yellow
	trafficLight.Change()
	if _, ok := trafficLight.State.(*YellowLightState); !ok {
		t.Error("Expected traffic light state to be YellowLightState after change from GreenLightState")
	}

	// Yellow to Red
	trafficLight.Change()
	if _, ok := trafficLight.State.(*RedLightState); !ok {
		t.Error("Expected traffic light state to be RedLightState after change from YellowLightState")
	}
}
