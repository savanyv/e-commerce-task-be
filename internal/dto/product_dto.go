package dtos

type ProductRequest struct {
	Name       string         `json:"name" validate:"required"`
	Type       string         `json:"type" validate:"required,oneof=single variant"`
	CategoryID int            `json:"category_id" validate:"required"`
	BrandID    int            `json:"brand_id" validate:"required"`
	Variants   []VariantInput `json:"variants" validate:"required_if=Type variant"`
}

type VariantInput struct {
	Name  string `json:"name" validate:"required"`
	Price int    `json:"price" validate:"required"`
	Stock int    `json:"stock" validate:"required"`
}

type ProductResponse struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
	CategoryID int `json:"category_id"`
	BrandID int `json:"brand_id"`
	Variants []VariantResponse `json:"variants"`
}

type VariantResponse struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Price int `json:"price"`
	Stock int `json:"stock"`
}
