package usecase

import (
	"errors"
	"strings"

	dtos "github.com/savanyv/e-commerce-task-be/internal/dto"
	"github.com/savanyv/e-commerce-task-be/internal/models"
	"github.com/savanyv/e-commerce-task-be/internal/repository"
)

type ProductUsecase interface {
	GetAllProducts() ([]*dtos.ProductResponse, error)
	GetProductByID(id int) (*dtos.ProductResponse, error)
	CreateProduct(req *dtos.ProductRequest) (*dtos.ProductResponse, error)
}

type productUsecase struct {
	repo repository.ProductRepository
}

func NewProductUsecase(repo repository.ProductRepository) ProductUsecase {
	return &productUsecase{
		repo: repo,
	}
}

func (u *productUsecase) GetAllProducts() ([]*dtos.ProductResponse, error) {
	products, err := u.repo.FindAll()
	if err != nil {
		return nil, errors.New("failed to get products")
	}

	var response []*dtos.ProductResponse
	for _, product := range products {
		response = append(response, &dtos.ProductResponse{
			ID:        product.ID,
			Name:      product.Name,
			Type:      product.Type,
			CategoryID: product.CategoryID,
			BrandID:   product.BrandID,
		})
	}

	return response, nil
}

func (u *productUsecase) GetProductByID(id int) (*dtos.ProductResponse, error) {
	product, err := u.repo.FindByID(id)
	if err != nil {
		return nil, errors.New("failed to get product")
	}

	response := &dtos.ProductResponse{
		ID:        product.ID,
		Name:      product.Name,
		Type:      product.Type,
		CategoryID: product.CategoryID,
		BrandID:   product.BrandID,
	}

	return response, nil
}

func (u *productUsecase) CreateProduct(req *dtos.ProductRequest) (*dtos.ProductResponse, error) {
	if strings.TrimSpace(req.Name) == "" {
		return nil, errors.New("name is required")
	}

	if req.Type != "single" && req.Type != "variant" {
		return nil, errors.New("type must be single or variant")
	}

	if req.CategoryID == 0 || req.BrandID == 0 {
		return nil, errors.New("category_id and brand_id are required")
	}

	product := &models.Product{
		Name:      req.Name,
		Type:      req.Type,
		CategoryID: req.CategoryID,
		BrandID:   req.BrandID,
	}

	if err := u.repo.Create(product); err != nil {
		return nil, errors.New("failed to create product")
	}

	response := &dtos.ProductResponse{
		ID:        product.ID,
		Name:      product.Name,
		Type:      product.Type,
		CategoryID: product.CategoryID,
		BrandID:   product.BrandID,
		Variants:  []dtos.VariantResponse{},
	}

	return response, nil
}
