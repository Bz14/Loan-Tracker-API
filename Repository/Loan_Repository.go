package repository

import (
	"context"
	domain "loan-tracker/Domain"
	infrastructure "loan-tracker/Infrastructure"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type LoanRepository struct {
	collection *mongo.Collection
	config   *infrastructure.Config
}

func NewLoanRepository(collection *mongo.Collection, config *infrastructure.Config) *LoanRepository {
	return &LoanRepository{
		collection: collection,
		config: config,
	}
}


func (lr *LoanRepository) CreateLoan(loan domain.Loan) error {
	context, _ := context.WithTimeout(context.Background(), time.Duration(lr.config.ContextTimeout) * time.Second)
	loan.ID = primitive.NewObjectID()
	_, err := lr.collection.InsertOne(context, loan)
	if err != nil {
		return err
	}
	return nil
}