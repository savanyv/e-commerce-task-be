package usecase

import (
	"errors"
	"strings"

	dtos "github.com/savanyv/e-commerce-task-be/internal/dto"
	"github.com/savanyv/e-commerce-task-be/internal/models"
	"github.com/savanyv/e-commerce-task-be/internal/repository"
)

type BrandUsecase interface {
	GetAllBrands() ([]*dtos.BrandResponse, error)
	CreateBrand(req *dtos.BrandRequest) (*dtos.BrandResponse, error)
}

type brandUsecase struct {
	repo repository.BrandRepository
}

func NewBrandUsecase(repo repository.BrandRepository) BrandUsecase {
	return &brandUsecase{
		repo: repo,
	}
}

func (u *brandUsecase) GetAllBrands() ([]*dtos.BrandResponse, error) {
	brands, err := u.repo.FindAll()
	if err != nil {
		return nil, errors.New("failed to get brands")
	}

	var response []*dtos.BrandResponse
	for _, brand := range brands {
		response = append(response, &dtos.BrandResponse{
			ID:   brand.ID,
			Name: brand.Name,
		})
	}

	return response, nil
}

func (u *brandUsecase) CreateBrand(req *dtos.BrandRequest) (*dtos.BrandResponse, error) {
	if strings.TrimSpace(req.Name) == "" {
		return nil, errors.New("name is required")
	}

	brand := &models.Brand{
		Name: req.Name,
	}

	if err := u.repo.Create(brand); err != nil {
		return nil, errors.New("failed to create brand")
	}

	response := &dtos.BrandResponse{
		ID:   brand.ID,
		Name: brand.Name,
	}

	return response, nil
}
