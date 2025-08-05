package helpers

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTClaim struct {
	UserID int `json:"user_id"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

type JWTService interface {
	GenerateToken(userID int, email string) (string, error)
	ValidateToken(token string) (*JWTClaim, error)
}

type jwtService struct {
	secretKey string
	issuer string
}

func NewJwtService() JWTService {
	return &jwtService{
		secretKey: "secret",
		issuer: "e-commerce",
	}
}

func (j *jwtService) GenerateToken(userID int, email string) (string, error) {
	claims := &JWTClaim{
		UserID: userID,
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: j.issuer,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(j.secretKey))
}

func (j *jwtService) ValidateToken(token string) (*JWTClaim, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &JWTClaim{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(j.secretKey), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := tokenClaims.Claims.(*JWTClaim)
	if !ok || !tokenClaims.Valid {
		return nil, jwt.ErrInvalidKey
	}

	return claims, nil
}
