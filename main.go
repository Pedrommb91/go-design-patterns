package main

import (
	"Pedrommb91/go-design-patterns/creational/builder"
	"Pedrommb91/go-design-patterns/creational/factory"
	"Pedrommb91/go-design-patterns/creational/prototype"
	"Pedrommb91/go-design-patterns/creational/singleton"
	"fmt"
)

func main() {
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
	_, err := factory.CreateProduct(factory.ProductA)
	if err != nil {
		panic(err)
	}

	// Abstract Factory
	factory1 := factory.GetAbstractFactory("1")
	productA := factory1.CreateProductA()
	productB := factory1.CreateProductB()
	fmt.Println(productA.UseA())
	fmt.Println(productB.UseB())

	factory2 := factory.GetAbstractFactory("2")
	productA = factory2.CreateProductA()
	productB = factory2.CreateProductB()
	fmt.Println(productA.UseA())
	fmt.Println(productB.UseB())

	// Prototype
	circle := &prototype.Circle{Radius: 10, Color: "red"}
	cloneCircle := circle.Clone().(*prototype.Circle)
	fmt.Printf("Original Circle: %+v, Clone Circle: %+v\n", circle, cloneCircle)

	rectangle := &prototype.Rectangle{Width: 20, Height: 10, Color: "blue"}
	cloneRectangle := rectangle.Clone().(*prototype.Rectangle)
	fmt.Printf("Original Rectangle: %+v, Clone Rectangle: %+v\n", rectangle, cloneRectangle)
}
