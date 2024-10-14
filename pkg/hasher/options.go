package hasher

// Option -.
type Option func(*Hasher)

// HasherCost -.
func HasherCost(cost int) Option {
	return func(c *Hasher) {
		c.Cost = cost
	}
}
