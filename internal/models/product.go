package models

type Product struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
	CategoryID int `json:"category_id"`
	BrandID int `json:"brand_id"`
	Variants []Variant `json:"variants"`
}

type Variant struct {
	ID        int    `json:"id"`
	ProductID int    `json:"product_id"`
	Name      string `json:"name"`
	Price     int    `json:"price"`
	Stock     int    `json:"stock"`
}
