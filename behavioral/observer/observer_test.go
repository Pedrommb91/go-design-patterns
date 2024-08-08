package observer

import (
	"testing"
)

func TestWeatherStation_RegisterDeregisterObserver(t *testing.T) {
	weatherStation := WeatherStation{}
	observer := &TemperatureDisplay{}

	weatherStation.RegisterObserver(observer)
	if len(weatherStation.observers) != 1 {
		t.Errorf("Expected 1 observer, got %d", len(weatherStation.observers))
	}

	weatherStation.DeregisterObserver(observer)
	if len(weatherStation.observers) != 0 {
		t.Errorf("Expected 0 observers, got %d", len(weatherStation.observers))
	}
}

func TestWeatherStation_NotifyObservers(t *testing.T) {
	weatherStation := WeatherStation{}
	tempDisplay := &TemperatureDisplay{}
	smartphoneApp := &SmartphoneApp{}

	weatherStation.RegisterObserver(tempDisplay)
	weatherStation.RegisterObserver(smartphoneApp)

	// Simulate temperature change
	weatherStation.SetTemperature(25.3)

	if tempDisplay.currentTemp != 25.3 {
		t.Errorf("Expected TemperatureDisplay to be updated to 25.3, got %.2f", tempDisplay.currentTemp)
	}

	if smartphoneApp.currentTemp != 25.3 {
		t.Errorf("Expected SmartphoneApp to be updated to 25.3, got %.2f", smartphoneApp.currentTemp)
	}
}

func TestWeatherStation_SetTemperature(t *testing.T) {
	weatherStation := WeatherStation{}
	weatherStation.SetTemperature(30.0)

	if weatherStation.temperature != 30.0 {
		t.Errorf("Expected temperature to be set to 30.0, got %.2f", weatherStation.temperature)
	}
}
