package domain

type LoanRequest struct {
	Amount float64 `json:"amount" validate:"required"`
	UserId string  `bson:"user_id" json:"user_id" validate:"required"`
}