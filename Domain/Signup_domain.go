package domain

type SignUpRequest struct {
	Email     string `json:"email" validate:"required"`
	User_Name string `json:"user_name" validate:"required"`
	Password  string `json:"password" validate:"required"`
}
