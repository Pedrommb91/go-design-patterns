package main

import "Pedrommb91/go-design-patterns/creational"

func main() {
	// Singleton
	_ = creational.GetInstance()

	// Simple Builder
	builder := creational.NewBuilder()
	builder.SetName("Product")
	builder.SetValue(100)
	_ = builder.Build()

}
