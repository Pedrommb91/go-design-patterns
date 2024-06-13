package decorator

// Coffee is the interface for all coffee types
type Coffee interface {
	Cost() float64
	Description() string
}

// SimpleCoffee is a basic coffee without any add-ons
type SimpleCoffee struct{}

func (s *SimpleCoffee) Cost() float64 {
	return 2.00
}

func (s *SimpleCoffee) Description() string {
	return "Simple Coffee"
}

// CoffeeDecorator is a base decorator that follows the Coffee interface
type CoffeeDecorator struct {
	coffee Coffee
}

func (d *CoffeeDecorator) Decorate(coffee Coffee) {
	d.coffee = coffee
}

func (d *CoffeeDecorator) Cost() float64 {
	return d.coffee.Cost()
}

func (d *CoffeeDecorator) Description() string {
	return d.coffee.Description()
}

// MilkDecorator is a concrete decorator that adds milk to the coffee
type MilkDecorator struct {
	CoffeeDecorator
}

func (m *MilkDecorator) Cost() float64 {
	return m.coffee.Cost() + 0.50
}

func (m *MilkDecorator) Description() string {
	return m.coffee.Description() + ", Milk"
}

// SugarDecorator is another concrete decorator that adds sugar to the coffee
type SugarDecorator struct {
	CoffeeDecorator
}

func (s *SugarDecorator) Cost() float64 {
	return s.coffee.Cost() + 0.25
}

func (s *SugarDecorator) Description() string {
	return s.coffee.Description() + ", Sugar"
}
