package jwtgenerator

import "time"

// Option -.
type Option func(*JwtTokenGenerator)

// TokenNbf -.
func TokenNbf(timeout time.Duration) Option {
	return func(c *JwtTokenGenerator) {
		c.Nbf = timeout
	}
}

// TokenExp -.
func TokenExp(timeout time.Duration) Option {
	return func(c *JwtTokenGenerator) {
		c.Exp = timeout
	}
}
