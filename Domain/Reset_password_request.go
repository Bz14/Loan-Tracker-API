package domain

type ResetPasswordRequest struct {
	Email string `json:"email" binding:"required"`
}