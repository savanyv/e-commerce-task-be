package models

type Order struct {
	ID int `json:"id"`
	UserID int `json:"user_id"`
	Items []OrderItem `json:"items"`
}

type OrderItem struct {
	ID int `json:"id"`
	OrderID int `json:"order_id"`
	VariantID int `json:"variant_id"`
	Quantity int `json:"quantity"`
	Price int `json:"price"`
}
