package creational

type Product struct {
	name  string
	value int
}

func (p Product) GetName() string {
	return p.name
}

func (p Product) GetValue() int {
	return p.value
}

type Builder struct {
	product Product
}

func NewBuilder() *Builder {
	return &Builder{
		product: Product{},
	}
}

func (b *Builder) SetName(name string) {
	b.product.name = name
}

func (b *Builder) SetValue(value int) {
	b.product.value = value
}

func (b *Builder) Build() Product {
	return b.product
}
