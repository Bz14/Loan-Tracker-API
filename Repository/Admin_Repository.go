package repository

import (
	"context"
	domain "loan-tracker/Domain"
	infrastructure "loan-tracker/Infrastructure"
	utils "loan-tracker/Utils"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AdminRepository struct {
	collection *mongo.Collection
	config   *infrastructure.Config
}

func NewAdminRepository(collection *mongo.Collection, config *infrastructure.Config) *AdminRepository {
	return &AdminRepository{
		collection: collection,
		config: config,
	}
}


func (ar *AdminRepository) GetAllUsers(pageNo, pageSize int64) ([]domain.User, error){
	var users []domain.User
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	
	options := utils.PaginationByPage(pageNo, pageSize)
	cursor, err := ar.collection.Find(ctx, bson.M{}, options)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var user domain.User
		cursor.Decode(&user)
		users = append(users, user)
	}
	return users, nil
}

func (ar *AdminRepository) DeleteUser(id string) error{
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	_, err := ar.collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return err
	}
	return nil
}