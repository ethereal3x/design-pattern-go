package new

type Product struct {
	Name  string
	Price float64
}

func (p *Product) GetName() string {
	return p.Name
}

func (p *Product) GetPrice() float64 {
	return p.Price
}

// NewProduct 创建并返回实例化后的实例，*Product 类型.
func NewProduct(name string, price float64) *Product {
	return &Product{
		Name:  name,
		Price: price,
	}
}
