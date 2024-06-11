package builder

type Product struct {
	name  string
	value int
	color string
	size  string
}

type Builder interface {
	SetName(name string) Builder
	SetValue(value int) Builder
	SetColor(color string) Builder
	SetSize(size string) Builder
	Build() Product
}

type ConcreteBuilder struct {
	product Product
}

type BuilderType int

const (
	Concrete BuilderType = iota
)

func NewBuilder(builderType BuilderType) Builder {
	switch builderType {
	case Concrete:
		return &ConcreteBuilder{
			product: Product{},
		}
	default:
		return nil
	}
}

func (b *ConcreteBuilder) SetName(name string) Builder {
	b.product.name = name
	return b
}

func (b *ConcreteBuilder) SetValue(value int) Builder {
	b.product.value = value
	return b
}

func (b *ConcreteBuilder) SetColor(color string) Builder {
	b.product.color = color
	return b
}

func (b *ConcreteBuilder) SetSize(size string) Builder {
	b.product.size = size
	return b
}

func (b *ConcreteBuilder) Build() Product {
	return b.product
}
