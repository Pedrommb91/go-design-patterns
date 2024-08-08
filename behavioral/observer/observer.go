package observer

import "fmt"

// WeatherStation acts as the Subject in the observer pattern.
type WeatherStation struct {
	temperature float64
	observers   []Observer
}

// RegisterObserver adds an Observer to the WeatherStation.
func (ws *WeatherStation) RegisterObserver(o Observer) {
	ws.observers = append(ws.observers, o)
}

// DeregisterObserver removes an Observer from the WeatherStation.
func (ws *WeatherStation) DeregisterObserver(o Observer) {
	for i, observer := range ws.observers {
		if observer == o {
			ws.observers = append(ws.observers[:i], ws.observers[i+1:]...)
			break
		}
	}
}

// NotifyObservers informs all registered Observers of a change.
func (ws *WeatherStation) NotifyObservers() {
	for _, observer := range ws.observers {
		observer.Update(ws.temperature)
	}
}

// SetTemperature changes the temperature and notifies Observers.
func (ws *WeatherStation) SetTemperature(temp float64) {
	ws.temperature = temp
	ws.NotifyObservers()
}

// Observer interface defines a method for updating the observer with the Subject's state.
type Observer interface {
	Update(temperature float64)
}

// TemperatureDisplay implements Observer and displays temperature changes.
type TemperatureDisplay struct {
	currentTemp float64
}

// Update method for TemperatureDisplay, called when WeatherStation changes temperature.
func (td *TemperatureDisplay) Update(temperature float64) {
	td.currentTemp = temperature
	td.Display()
}

// Display prints the current temperature to Stdout.
func (td *TemperatureDisplay) Display() {
	fmt.Printf("[Display] Current temperature is: %.2f°C\n", td.currentTemp)
}

// SmartphoneApp implements Observer and could represent a user's app that receives updates.
type SmartphoneApp struct {
	currentTemp float64
}

// Update method for SmartphoneApp, called when WeatherStation changes temperature.
func (sa *SmartphoneApp) Update(temperature float64) {
	sa.currentTemp = temperature
	sa.ShowNotification()
}

// ShowNotification simulates sending a temperature update notification to a smartphone.
func (sa *SmartphoneApp) ShowNotification() {
	fmt.Printf("[SmartPhoneApp] Temperature update: %.2f°C\n", sa.currentTemp)
}
