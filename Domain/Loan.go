package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Loan struct {
	ID     primitive.ObjectID `bson:"_id,omitempity" json:"id" `
	Amount float64            `bson:"amount" json:"amount" validate:"required"`
	UserId  primitive.ObjectID             `bson:"user_id" json:"user_id" validate:"required"`
	LoanStatus string         `bson:"loan_status" json:"loan_status"`
	Created_at time.Time	  `bson:"created_at" json:"created_at"`
}


type LoanUseCaseInterface interface {
	CreateLoan(loan Loan, user_id string) error
	CheckLoanStatus(id string, user_id string) (string, error)
	GetAllLoans(status string, order string, user_id string) ([]Loan, error)
}

type LoanRepositoryInterface interface {
	CreateLoan(loan Loan) error
	GetAllLoans(status string, order string) ([]Loan, error)
	FindLoanByID(id string)(Loan , error)
}