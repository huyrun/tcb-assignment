package auth

import (
	"github.com/google/uuid"
)

type Service interface {
	getSecretJWT() string
}

type UID struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
}

type service struct {
	secretJWT string
}

func NewAuthService(
	secretJWT string,
) *service {
	return &service{
		secretJWT: secretJWT,
	}
}

func (s *service) getSecretJWT() string {
	return s.secretJWT
}
