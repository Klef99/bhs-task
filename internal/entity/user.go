package entity

type User struct {
	Id       int64
	Username string
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
