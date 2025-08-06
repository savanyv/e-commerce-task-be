package usecase

import (
	"errors"
	"strings"

	dtos "github.com/savanyv/e-commerce-task-be/internal/dto"
	"github.com/savanyv/e-commerce-task-be/internal/models"
	"github.com/savanyv/e-commerce-task-be/internal/repository"
)

type CategoryUsecase interface {
	GetAllCategories() ([]*dtos.CategoryResponse, error)
	CreateCategory(req *dtos.CategoryRequest) (*dtos.CategoryResponse, error)
}

type categoryUsecase struct {
	repo repository.CategoryRepository
}

func NewCategoryUsecase(repo repository.CategoryRepository) CategoryUsecase {
	return &categoryUsecase{
		repo: repo,
	}
}

func (u *categoryUsecase) GetAllCategories() ([]*dtos.CategoryResponse, error) {
	categories, err := u.repo.FindAll()
	if err != nil {
		return nil, errors.New("failed to get categories")
	}

	var response []*dtos.CategoryResponse
	for _, category := range categories {
		response = append(response, &dtos.CategoryResponse{
			ID:   category.ID,
			Name: category.Name,
		})
	}

	return response, nil
}

func (u *categoryUsecase) CreateCategory(req *dtos.CategoryRequest) (*dtos.CategoryResponse, error) {
	if strings.TrimSpace(req.Name) == "" {
		return nil, errors.New("name is required")
	}

	category := &models.Category{
		Name: req.Name,
	}

	if err := u.repo.Create(category); err != nil {
		return nil, errors.New("failed to create category")
	}

	response := &dtos.CategoryResponse{
		ID:   category.ID,
		Name: category.Name,
	}

	return response, nil
}
