package entity

import "fmt"

type User struct {
	Id       int64
	Username string
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (c Credentials) Validate() error {
	if c.Username == "" || c.Password == "" {
		return fmt.Errorf("username or password are invalid")
	}
	return nil
}
