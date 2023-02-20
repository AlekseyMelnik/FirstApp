package dto

type CreateUserDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type UpdateUserDTO struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
