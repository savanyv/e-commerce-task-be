package usecase

import (
	"errors"

	dtos "github.com/savanyv/e-commerce-task-be/internal/dto"
	"github.com/savanyv/e-commerce-task-be/internal/models"
	"github.com/savanyv/e-commerce-task-be/internal/repository"
)

type OrderUsecase interface {
	Checkout(req *dtos.OrderRequest) (*dtos.OrderResponse, error)
}

type orderUsecase struct {
	repo repository.OrderRepository
	variantRepo repository.VariantRepository
}

func NewOrderUsecase(repo repository.OrderRepository, variantRepo repository.VariantRepository) OrderUsecase {
	return &orderUsecase{
		repo: repo,
		variantRepo: variantRepo,
	}
}

func (u *orderUsecase) Checkout(req *dtos.OrderRequest) (*dtos.OrderResponse, error) {
	if len(req.Items) == 0 {
		return nil, errors.New("no items to checkout")
	}

	var orderItems []models.OrderItem
	var responseItems []dtos.OrderItemResponse
	var totalPrice int

	for _, item := range req.Items {
		variant, err := u.variantRepo.FindByID(item.VariantID)
		if err != nil {
			return nil, errors.New("variant not found")
		}

		if variant.Stock < item.Quantity {
			return nil, errors.New("not enough stock")
		}

		orderItems = append(orderItems, models.OrderItem{
			VariantID: item.VariantID,
			Quantity:  item.Quantity,
			Price:     variant.Price,
		})

		responseItems = append(responseItems, dtos.OrderItemResponse{
			VariantID: item.VariantID,
			Quantity:  item.Quantity,
			Price:     variant.Price,
		})

		totalPrice += variant.Price * item.Quantity
	}

	order := &models.Order{
		UserID: req.UserID,
		Items:  orderItems,
	}

	if err := u.repo.Create(order); err != nil {
		return nil, errors.New("failed to create order")
	}

	response := &dtos.OrderResponse{
		ID:         order.ID,
		TotalPrice: totalPrice,
		Items:      responseItems,
	}

	return response, nil
}

