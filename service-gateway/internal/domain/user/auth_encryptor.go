package user

import (
	"golang.org/x/crypto/bcrypt"
)

// Interface encryptor interface
type Encryptor interface {
	Digest(password string) (string, error)
	Compare(hashedPassword string, password string) error
}

type BcryptEncryptor struct {
	cost int
}

func NewEncryptor(cost int) *BcryptEncryptor {
	if cost == 0 {
		cost = bcrypt.DefaultCost
	}

	return &BcryptEncryptor{
		cost: cost,
	}
}

// Digest generate encrypted password
func (b *BcryptEncryptor) Digest(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword), err
}

// Compare check hashed password
func (b *BcryptEncryptor) Compare(hashedPassword string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
