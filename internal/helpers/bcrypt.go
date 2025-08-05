package helpers

import "golang.org/x/crypto/bcrypt"

type BcryptService interface {
	HashPassword(password string) (string, error)
	ComparePassword(password, hashedPassword string) error
}

type bcryptService struct {
	cost int
}

func NewBcryptService() BcryptService {
	return &bcryptService{
		cost: bcrypt.DefaultCost,
	}
}

func (b *bcryptService) HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), b.cost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func (b *bcryptService) ComparePassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
