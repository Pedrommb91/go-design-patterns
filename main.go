package main

import (
	"Pedrommb91/go-design-patterns/creational/builder"
	"Pedrommb91/go-design-patterns/creational/factory"
	"Pedrommb91/go-design-patterns/creational/prototype"
	"Pedrommb91/go-design-patterns/creational/singleton"

	"Pedrommb91/go-design-patterns/structural/adapter"
	"Pedrommb91/go-design-patterns/structural/bridge"
	"Pedrommb91/go-design-patterns/structural/composite"
	"Pedrommb91/go-design-patterns/structural/decorator"
	"Pedrommb91/go-design-patterns/structural/facade"
	"Pedrommb91/go-design-patterns/structural/flyweight"
	"Pedrommb91/go-design-patterns/structural/proxy"

	"Pedrommb91/go-design-patterns/behavioral/chain"
	"Pedrommb91/go-design-patterns/behavioral/command"
	"Pedrommb91/go-design-patterns/behavioral/interpreter"
	"Pedrommb91/go-design-patterns/behavioral/iterator"
	"Pedrommb91/go-design-patterns/behavioral/mediator"
	"Pedrommb91/go-design-patterns/behavioral/memento"

	"fmt"
)

func main() {
	// ++++++++++++++++++++++++++++++ Creational ++++++++++++++++++++++++++++++
	// Singleton
	_ = singleton.GetInstance()

	// Simple Builder
	_ = builder.NewSimpleBuilder().
		SetName("Product").
		SetValue(100).
		Build()

	// Builder
	_ = builder.NewBuilder(builder.Default).
		Build()

	_ = builder.NewBuilder(builder.Custom).
		SetName("Product").
		SetValue(100).
		Build()

	// Factory method
	vehicle, err := factory.CreateVehicle(factory.CarType)
	if err != nil {
		panic(err)
	}
	fmt.Println(vehicle.Drive())

	// Abstract Factory
	factory1 := factory.GetFurnitureFactory("modern")
	if factory1 != nil {
		chair := factory1.CreateChair()
		table := factory1.CreateTable()
		fmt.Println(chair.SitOn())
		fmt.Println(table.PlaceItems())
	} else {
		fmt.Println("Factory not found for style: modern")
	}

	factory2 := factory.GetFurnitureFactory("victorian")
	if factory2 != nil {
		chair := factory2.CreateChair()
		table := factory2.CreateTable()
		fmt.Println(chair.SitOn())
		fmt.Println(table.PlaceItems())
	} else {
		fmt.Println("Factory not found for style: victorian")
	}

	// Prototype
	circle := &prototype.Circle{Radius: 10, Color: "red"}
	cloneCircle := circle.Clone().(*prototype.Circle)
	fmt.Printf("Original Circle: %+v, Clone Circle: %+v\n", circle, cloneCircle)

	rectangle := &prototype.Rectangle{Width: 20, Height: 10, Color: "blue"}
	cloneRectangle := rectangle.Clone().(*prototype.Rectangle)
	fmt.Printf("Original Rectangle: %+v, Clone Rectangle: %+v\n", rectangle, cloneRectangle)

	// ++++++++++++++++++++++++++++++ Structural ++++++++++++++++++++++++++++++
	// Adapter
	usbDevice := &adapter.USBDevice{}
	adapter.ConnectDeviceToComputer(usbDevice)

	hdmiDevice := &adapter.HDMIDevice{}
	hdmiToUSBAdapter := &adapter.HDMIToUSBAdapter{HDMIDevice: hdmiDevice}
	adapter.ConnectDeviceToComputer(hdmiToUSBAdapter)

	// Bridge
	tv := &bridge.TV{}
	radio := &bridge.Radio{}

	remote := &bridge.RemoteControl{Device: tv}
	remote.TurnOn()
	remote.TurnOff()

	advancedRemote := &bridge.AdvancedRemoteControl{RemoteControl: bridge.RemoteControl{Device: radio}}
	advancedRemote.TurnOn()
	advancedRemote.Mute()
	advancedRemote.TurnOff()

	// Composite
	file1 := &composite.File{Name: "File 1"}
	file2 := &composite.File{Name: "File 2"}
	file3 := &composite.File{Name: "File 3"}

	dir1 := &composite.Directory{Name: "Directory 1"}
	dir1.Add(file1)
	dir1.Add(file2)

	dir2 := &composite.Directory{Name: "Directory 2"}
	dir2.Add(dir1)
	dir2.Add(file3)

	fmt.Println(file1.Operation())
	fmt.Println(dir1.Operation())
	fmt.Println(dir2.Operation())

	// Decorator
	// Create a simple coffee
	coffee := &decorator.SimpleCoffee{}

	// Decorate it with Milk
	milkCoffee := &decorator.MilkDecorator{}
	milkCoffee.Decorate(coffee)
	fmt.Printf("%s: $%.2f\n", milkCoffee.Description(), milkCoffee.Cost()) // Output: Simple Coffee, Milk: $2.50

	// Decorate it with Sugar
	sugarCoffee := &decorator.SugarDecorator{}
	sugarCoffee.Decorate(milkCoffee)
	fmt.Printf("%s: $%.2f\n", sugarCoffee.Description(), sugarCoffee.Cost()) // Output: Simple Coffee, Milk, Sugar: $2.75

	// Facade
	amplifier := &facade.Amplifier{}
	dvdPlayer := &facade.DvdPlayer{}
	projector := &facade.Projector{}
	theaterLights := &facade.TheaterLights{}
	screen := &facade.Screen{}

	homeTheater := facade.NewHomeTheaterFacade(amplifier, dvdPlayer, projector, theaterLights, screen)
	homeTheater.WatchMovie("Pulp Fiction")
	homeTheater.EndMovie()

	// flyweight
	forest := flyweight.NewForest()

	// Planting 5 pine trees in the forest at different locations
	forest.PlantTree(1, 2, 5, 2, "Pine", "Green", "Pine Texture")
	forest.PlantTree(2, 3, 5, 2, "Pine", "Green", "Pine Texture")
	forest.PlantTree(3, 4, 5, 2, "Pine", "Green", "Pine Texture")
	forest.PlantTree(4, 5, 5, 2, "Pine", "Green", "Pine Texture")
	forest.PlantTree(5, 6, 5, 2, "Pine", "Green", "Pine Texture")

	// Planting 3 oak trees in the forest at different locations
	forest.PlantTree(6, 7, 10, 5, "Oak", "Brown", "Oak Texture")
	forest.PlantTree(7, 8, 10, 5, "Oak", "Brown", "Oak Texture")
	forest.PlantTree(8, 9, 10, 5, "Oak", "Brown", "Oak Texture")

	// Assuming we have a canvas to draw on
	canvas := &flyweight.Canvas{}

	// Drawing all trees in the forest
	forest.Draw(canvas)

	// Proxy
	door := &proxy.LabDoor{}
	securedDoor := proxy.NewSecurity(door, "secret")

	fmt.Println(securedDoor.Open("invalid")) // Output: Access denied
	fmt.Println(securedDoor.Open("secret"))  // Output: Lab door opened
	fmt.Println(securedDoor.Close())         // Output: Lab door closed

	// ++++++++++++++++++++++++++++++ Behavioral ++++++++++++++++++++++++++++++
	// Chain of Responsibility
	approver := &chain.BaseApprover{}
	manager := &chain.Manager{}
	director := &chain.Director{}
	vicePresident := &chain.VicePresident{}
	ceo := &chain.CEO{}

	// Set up the chain of responsibility
	approver.SetNext(manager).SetNext(director).SetNext(vicePresident).SetNext(ceo)

	// Various requests
	expenses := []float64{500, 1500, 5500, 10500}
	for _, amount := range expenses {
		approved := approver.ApproveRequest(amount)
		if approved {
			fmt.Printf("Request for $%.2f has been approved.\n", amount)
		} else {
			fmt.Printf("Request for $%.2f has been denied.\n", amount)
		}
	}

	// Command
	light := &command.Light{}
	switcher := command.NewSwitch()
	switcher.StoreCommand("on", &command.TurnOnCommand{Light: light})
	switcher.StoreCommand("off", &command.TurnOffCommand{Light: light})

	// Execute the commands
	switcher.ExecuteCommand("on")
	switcher.ExecuteCommand("off")

	// Interpreter
	ctx := &interpreter.Context{
		VariableValues: map[string]int{"x": 5, "y": 10},
	}
	expression := &interpreter.AddExpression{
		Left:  &interpreter.VariableExpression{Name: "x"},
		Right: &interpreter.VariableExpression{Name: "y"},
	}
	result := expression.Interpret(ctx)
	fmt.Println(result) // Output: 15

	// Iterator
	bookShelf := &iterator.BookShelf{
		Books: []iterator.Book{
			{Title: "Design Patterns", Author: "Erich Gamma"},
			{Title: "Clean Code", Author: "Robert C. Martin"},
			{Title: "Refactoring", Author: "Martin Fowler"},
		},
	}

	it := bookShelf.CreateIterator()
	for it.HasNext() {
		book := it.Next()
		fmt.Printf("Book: %s, Author: %s\n", book.Title, book.Author)
	}

	// Mediator
	controlTower := mediator.NewControlTower()
	passengerPlane1 := mediator.NewPassengerPlane(controlTower, 1)
	passengerPlane2 := mediator.NewPassengerPlane(controlTower, 2)

	if passengerPlane1.Land() {
		fmt.Println("Passenger Plane 1 has landed.")
	} else {
		fmt.Println("Passenger Plane 1 is waiting to land.")
	}

	if passengerPlane2.Land() {
		fmt.Println("Passenger Plane 2 has landed.")
	} else {
		fmt.Println("Passenger Plane 2 is waiting to land.")
	}

	passengerPlane1.Takeoff()
	if passengerPlane2.Land() {
		fmt.Println("Passenger Plane 2 has landed.")
	}

	// Memento
	editor := memento.Editor{}
	history := memento.History{}

	editor.SetContent("First State")
	history.Push(editor.CreateState())

	editor.SetContent("Second State")
	history.Push(editor.CreateState())

	editor.SetContent("Third State")

	fmt.Println(editor.GetContent())

	// Restore to the second state
	secondState := history.Pop()
	if secondState != nil {
		editor.Restore(secondState)
	}
	fmt.Println(editor.GetContent()) // Output: Second State

	// Restore to the first state
	firstState := history.Pop()
	if firstState != nil {
		editor.Restore(firstState)
	}
	fmt.Println(editor.GetContent()) // Output: First State

}
