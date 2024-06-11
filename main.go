package main

import (
	"Pedrommb91/go-design-patterns/creational/builder"
	"Pedrommb91/go-design-patterns/creational/factory"
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
	// Use ConcreteFactory1
	factory1 := factory.GetAbstractFactory("1")
	productA := factory1.CreateProductA()
	productB := factory1.CreateProductB()
	fmt.Println(productA.UseA())
	fmt.Println(productB.UseB())

	// Use ConcreteFactory2
	factory2 := factory.GetAbstractFactory("2")
	productA = factory2.CreateProductA()
	productB = factory2.CreateProductB()
	fmt.Println(productA.UseA())
	fmt.Println(productB.UseB())
}
