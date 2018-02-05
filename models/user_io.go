package models

type LoginForm struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type RegisterForm struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type UserOut struct {
	Id    int    `json:"id"`
	Phone string `json:"phone"`
	Name  string `json:"name"`
}
