package hasher

import (
	"golang.org/x/crypto/bcrypt"
)

const (
	_defaultCost = bcrypt.DefaultCost
)

type Interface interface {
	HashPassword(password string) ([]byte, error)
	CompareHashAndPassword(password_hash, password string) error
}

type Hasher struct {
	Cost int
}

func NewHasher(opts ...Option) *Hasher {
	hs := &Hasher{Cost: _defaultCost}
	// Custom options
	for _, opt := range opts {
		opt(hs)
	}
	return hs
}

var _ Interface = (*Hasher)(nil)

func (h *Hasher) HashPassword(password string) ([]byte, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), h.Cost)
	if err != nil {
		return []byte{}, err
	}
	return hashedBytes, nil
}

func (h *Hasher) CompareHashAndPassword(passwordHash, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	return err
}
