package domain

type ResetPassword struct {
	NewPassword string `json:"new_password" validate:"required"`
}