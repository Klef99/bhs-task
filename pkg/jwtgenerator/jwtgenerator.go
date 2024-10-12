package jwtgenerator

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	_defaultNbf = time.Nanosecond
	_defaultExp = time.Hour
)

type Interface interface {
	GenerateToken(username string) (string, error)
	ValidateToken(tokenString string) (string, error)
}

type JwtTokenGenerator struct {
	secret string
	Nbf    time.Duration
	Exp    time.Duration
}

var _ Interface = (*JwtTokenGenerator)(nil)

func (jtg *JwtTokenGenerator) GenerateToken(username string) (string, error) {
	now := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name": username,
		"nbf":  now.Add(jtg.Nbf).Unix(),
		"exp":  now.Add(jtg.Exp).Unix(),
		"iat":  now.Unix(),
	})
	tokenString, err := token.SignedString([]byte(jtg.secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// ValidateToken - return name of the token owner, if token is valid
func (jtg *JwtTokenGenerator) ValidateToken(tokenString string) (string, error) {
	tokenFromString, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return "", fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(jtg.secret), nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := tokenFromString.Claims.(jwt.MapClaims); ok {
		return claims["name"].(string), nil
	} else {
		return "", err
	}
}

// New -.
func New(secret string, opts ...Option) (*JwtTokenGenerator, error) {
	jtg := &JwtTokenGenerator{
		secret: secret,
		Nbf:    _defaultNbf,
		Exp:    _defaultExp,
	}

	// Custom options
	for _, opt := range opts {
		opt(jtg)
	}
	return jtg, nil
}
