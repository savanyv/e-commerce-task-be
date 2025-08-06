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
	variantRepo repository.VariantRepository
}

func NewProductUsecase(repo repository.ProductRepository, variantRepo repository.VariantRepository) ProductUsecase {
	return &productUsecase{
		repo: repo,
		variantRepo: variantRepo,
	}
}

func (u *productUsecase) GetAllProducts() ([]*dtos.ProductResponse, error) {
	products, err := u.repo.FindAll()
	if err != nil {
		return nil, errors.New("failed to get products")
	}

	var response []*dtos.ProductResponse
	for _, product := range products {
		variants, err := u.variantRepo.FindByProductID(product.ID)
		if err != nil {
			return nil, errors.New("failed to get variants")
		}

		var variantReponses []dtos.VariantResponse
		for _, v := range variants {
			variantReponses = append(variantReponses, dtos.VariantResponse{
				ID: v.ID,
				Name: v.Name,
				Price: v.Price,
				Stock: v.Stock,
			})
		}

		response = append(response, &dtos.ProductResponse{
			ID:        product.ID,
			Name:      product.Name,
			Type:      product.Type,
			CategoryID: product.CategoryID,
			BrandID:   product.BrandID,
			Variants:  variantReponses,
		})
	}

	return response, nil
}

func (u *productUsecase) GetProductByID(id int) (*dtos.ProductResponse, error) {
	product, err := u.repo.FindByID(id)
	if err != nil {
		return nil, errors.New("failed to get product")
	}

	variantModels, err := u.variantRepo.FindByProductID(product.ID)
	if err != nil {
		return nil, errors.New("failed to get variants")
	}

	var variantResponses []dtos.VariantResponse
	for _, v := range variantModels {
		variantResponses = append(variantResponses, dtos.VariantResponse{
			ID: v.ID,
			Name: v.Name,
			Price: v.Price,
			Stock: v.Stock,
		})
	}

	response := &dtos.ProductResponse{
		ID:        product.ID,
		Name:      product.Name,
		Type:      product.Type,
		CategoryID: product.CategoryID,
		BrandID:   product.BrandID,
		Variants:  variantResponses,
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
		Name:       req.Name,
		Type:       req.Type,
		CategoryID: req.CategoryID,
		BrandID:    req.BrandID,
	}

	if err := u.repo.Create(product); err != nil {
		return nil, errors.New("failed to create product")
	}

	var variantResponses []dtos.VariantResponse
	if len(req.Variants) > 0 {
		var variants []*models.Variant
		for _, v := range req.Variants {
			variants = append(variants, &models.Variant{
				ProductID: product.ID,
				Name:      v.Name,
				Price:     v.Price,
				Stock:     v.Stock,
			})
		}

		if err := u.variantRepo.CreateBulk(product.ID, variants); err != nil {
			return nil, errors.New("failed to create variants")
		}

		for _, v := range variants {
			variantResponses = append(variantResponses, dtos.VariantResponse{
				ID:    v.ID,
				Name:  v.Name,
				Price: v.Price,
				Stock: v.Stock,
			})
		}
	}

	response := &dtos.ProductResponse{
		ID:         product.ID,
		Name:       product.Name,
		Type:       product.Type,
		CategoryID: product.CategoryID,
		BrandID:    product.BrandID,
		Variants:   variantResponses,
	}

	return response, nil
}

