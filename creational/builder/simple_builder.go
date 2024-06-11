package builder

type SimpleProduct struct {
	name  string
	value int
}

func (p SimpleProduct) GetName() string {
	return p.name
}

func (p SimpleProduct) GetValue() int {
	return p.value
}

type SimpleBuilder struct {
	product SimpleProduct
}

func NewSimpleBuilder() *SimpleBuilder {
	return &SimpleBuilder{
		product: SimpleProduct{},
	}
}

func (b *SimpleBuilder) SetName(name string) *SimpleBuilder {
	b.product.name = name
	return b
}

func (b *SimpleBuilder) SetValue(value int) *SimpleBuilder {
	b.product.value = value
	return b
}

func (b *SimpleBuilder) Build() SimpleProduct {
	return b.product
}
