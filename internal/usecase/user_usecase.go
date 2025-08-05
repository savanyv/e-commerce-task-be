package usecase

import (
	"errors"

	dtos "github.com/savanyv/e-commerce-task-be/internal/dto"
	"github.com/savanyv/e-commerce-task-be/internal/helpers"
	"github.com/savanyv/e-commerce-task-be/internal/models"
	"github.com/savanyv/e-commerce-task-be/internal/repository"
)

type UserUsecase interface {
	Register(req *dtos.RegisterRequest) (*dtos.UserResponse, error)
	Login(req *dtos.LoginRequest) (*dtos.UserResponse, error)
}

type userUsecase struct {
	repo repository.UserRepository
	bcrypt helpers.BcryptService
	jwt helpers.JWTService
}

func NewUserUsecase(repo repository.UserRepository) UserUsecase {
	return &userUsecase{
		repo: repo,
		bcrypt: helpers.NewBcryptService(),
		jwt: helpers.NewJwtService(),
	}
}

func (u *userUsecase) Register(req *dtos.RegisterRequest) (*dtos.UserResponse, error) {
	existingUser, err := u.repo.FindByEmail(req.Email)
	if err == nil && existingUser != nil {
		return nil, errors.New("user already exists")
	}

	hashPassword, err := u.bcrypt.HashPassword(req.Password)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	user := &models.User{
		Email:    req.Email,
		Password: hashPassword,
	}

	if err := u.repo.Create(user); err != nil {
		return nil, errors.New("failed to create user")
	}

	response := &dtos.UserResponse{
		ID:    user.ID,
		Email: user.Email,
	}

	return response, nil
}

func (u *userUsecase) Login(req *dtos.LoginRequest) (*dtos.UserResponse, error) {
	user, err := u.repo.FindByEmail(req.Email)
	if err != nil {
		return nil, errors.New("user not found")
	}

	if err := u.bcrypt.ComparePassword(req.Password, user.Password); err != nil {
		return nil, errors.New("invalid password")
	}

	token, err := u.jwt.GenerateToken(user.ID, user.Email)
	if err != nil {
		return nil, errors.New("failed to generate token")
	}

	response := &dtos.UserResponse{
		ID:       user.ID,
		Email:    user.Email,
		Token: token,
	}

	return response, nil
}
