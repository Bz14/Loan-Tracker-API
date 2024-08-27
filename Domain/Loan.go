package domain

import "go.mongodb.org/mongo-driver/bson/primitive"

type Loan struct {
	ID     primitive.ObjectID `bson:"_id,omitempity" json:"id" `
	Amount float64            `bson:"amount" json:"amount" validate:"required"`
	UserId string             `bson:"user_id" json:"user_id" validate:"required"`
	LoanStatus string         `bson:"loan_status" json:"loan_status"`
}


type LoanUseCaseInterface interface {
	CreateLoan(loan Loan, user_id string) error
}

type LoanRepositoryInterface interface {
	CreateLoan(loan Loan) error
}