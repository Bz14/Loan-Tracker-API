package repository

import (
	"context"
	domain "loan-tracker/Domain"
	infrastructure "loan-tracker/Infrastructure"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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


func (lr *LoanRepository) FindLoanByID(id string)(domain.Loan , error){
	var Loan domain.Loan
	context, _ := context.WithTimeout(context.Background(), time.Duration(lr.config.ContextTimeout) * time.Second)
	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objectId}
	err := lr.collection.FindOne(context, filter).Decode(&Loan)
	if err != nil {
		return Loan, err
	}
	return Loan, nil
}

func (lr *LoanRepository) GetAllLoans(status string, order string) ([]domain.Loan, error){
	var loans []domain.Loan
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	findOptions := options.Find()
	if order != ""{
		if strings.ToLower(order) == "asc"{
			findOptions.SetSort(bson.D{{Key: "createdAt", Value: -1}})
		} else if strings.ToLower(order) == "desc"{
			findOptions.SetSort(bson.D{{Key: "createdAt", Value: 1}})
		}
	}
	filter := bson.M{}
	if status != ""{
		filter["loan_status"] = status
	}
	cursor, err := lr.collection.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var loan domain.Loan
		cursor.Decode(&loan)
		loans = append(loans, loan)
	}
	return loans, nil
}