package dtos

type OrderItemRequest struct {
	VariantID int `json:"variant_id" validate:"required"`
	Quantity  int `json:"quantity" validate:"required"`
}

type OrderRequest struct {
	UserID int `json:"user_id" validate:"required"`
	Items []OrderItemRequest `json:"items" validate:"required,dive"`
}

type OrderResponse struct {
	ID int `json:"id"`
	Items []OrderItemResponse `json:"items"`
	TotalPrice int `json:"total_price"`
}

type OrderItemResponse struct {
	VariantID int `json:"variant_id"`
	Quantity  int `json:"quantity"`
	Price     int `json:"price"`
}
