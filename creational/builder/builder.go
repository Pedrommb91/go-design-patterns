package builder

type Product struct {
	name  string
	value int
}

type Builder interface {
	SetName(name string) Builder
	SetValue(value int) Builder
	Build() Product
}

type DefaultBuilder struct {
	product Product
}

type CustomBuilder struct {
	product Product
}

type BuilderType int

const (
	Default BuilderType = iota
	Custom
)

func NewBuilder(builderType BuilderType) Builder {
	switch builderType {
	case Default:
		return &DefaultBuilder{
			product: Product{
				name:  "Default",
				value: 0,
			},
		}
	case Custom:
		return &CustomBuilder{
			product: Product{},
		}
	default:
		return nil
	}
}

func (b *DefaultBuilder) SetName(name string) Builder {
	b.product.name = name
	return b
}

func (b *DefaultBuilder) SetValue(value int) Builder {
	b.product.value = value
	return b
}

func (b *DefaultBuilder) Build() Product {
	return b.product
}

func (b *CustomBuilder) SetName(name string) Builder {
	b.product.name = name
	return b
}

func (b *CustomBuilder) SetValue(value int) Builder {
	b.product.value = value
	return b
}

func (b *CustomBuilder) Build() Product {
	return b.product
}
