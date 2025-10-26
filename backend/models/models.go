package models

type Product struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type PriceTag struct {
	ID            string  `json:"id"`
	ProductID     string  `json:"productId"`
	ProductName   string  `json:"productName"`
	Price         float64 `json:"price"`
	PriceTagImage string  `json:"priceTagImage"` // URL PNG ценника
}
