package dto

type LoginUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreateUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthUser struct {
	ID uint
}
