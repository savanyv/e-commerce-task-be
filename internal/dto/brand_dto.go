package dtos

type BrandRequest struct {
	Name string `json:"name" validate:"required"`
}

type BrandResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
