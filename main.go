package main

import (
	"Pedrommb91/go-design-patterns/creational/builder"
	"Pedrommb91/go-design-patterns/creational/factory"
	"Pedrommb91/go-design-patterns/creational/prototype"
	"Pedrommb91/go-design-patterns/creational/singleton"

	"Pedrommb91/go-design-patterns/structural/adapter"
	"Pedrommb91/go-design-patterns/structural/bridge"
	"Pedrommb91/go-design-patterns/structural/composite"
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
}
