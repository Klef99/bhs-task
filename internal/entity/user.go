package entity

type User struct {
	Id      int
	Usename string
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
