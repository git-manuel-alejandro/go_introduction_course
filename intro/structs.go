package intro

type Product struct {
	ID    uint    `json:"id"`
	Name  string  `json:"name"`
	Count uint16  `json:"count"`
	Type  string  `json:"type"`
	Price float64 `json:"price"`
}

func (p Product) TotalPrice() float64 {

	return float64(p.Count)

}

type Carro struct {
	ID       uint      `json:"id"`
	UserID   uint      `json:"user_id"`
	Products []Product `json:"products"`
}

func (c *Carro) AddProducts(products ...Product) {

	c.Products = append(c.Products, products...)

}

func NewCar(userId uint) Carro {
	return Carro{UserID: userId}
}
