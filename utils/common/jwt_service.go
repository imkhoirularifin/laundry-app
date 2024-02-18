package common

import (
	"laundry-app/config"
	"laundry-app/entity"
	"laundry-app/entity/dto"
	modelutil "laundry-app/utils/model_util"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtService interface {
	GenerateToken(payload entity.Employee) (dto.AuthResponse, error)
	VerifyToken(token string) (jwt.MapClaims, error)
	// TODO: add refresh token
	// RefreshToken(oldTokenString string) (dto.AuthResponse, error)
}

type jwtService struct {
	cfg config.JWTConfig
}

func (j *jwtService) GenerateToken(payload entity.Employee) (dto.AuthResponse, error) {
	// claims is jwt payload
	claims := modelutil.JwtTokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    payload.Name,
			Subject:   payload.Id,
			IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
			ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(j.cfg.Lifetime)),
		},
	}

	// sign token
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(j.cfg.SecretKey))
	if err != nil {
		return dto.AuthResponse{}, err
	}

	return dto.AuthResponse{Token: token}, nil
}

func (j *jwtService) VerifyToken(token string) (jwt.MapClaims, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(j.cfg.SecretKey), nil
	})

	return claims, err
}

func NewJwtToken(cfg config.JWTConfig) JwtService {
	return &jwtService{cfg: cfg}
}
