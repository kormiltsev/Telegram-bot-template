package product

type Sku struct {
	Title       string
	Description string
	Price       float64
}

type Catalog struct {
	Products []Sku
}
