package repository

import (
	"context"
	domain "loan-tracker/Domain"
	infrastructure "loan-tracker/Infrastructure"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	collection *mongo.Collection
	config   *infrastructure.Config
}

func NewUserRepository(collection *mongo.Collection, config *infrastructure.Config) *UserRepository {
	return &UserRepository{
		collection: collection,
		config: config,
	}
}

func (ur *UserRepository) FindUserByEmail(email string) (domain.User, error) {
	var user domain.User
	context, _ := context.WithTimeout(context.Background(), time.Duration(ur.config.ContextTimeout) * time.Second)
	filter := bson.M{"email": email}
	err := ur.collection.FindOne(context, filter).Decode(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (ur *UserRepository) FindUserByUserName(username string) (domain.User, error) {
	var user domain.User
	context, _ := context.WithTimeout(context.Background(), time.Duration(ur.config.ContextTimeout) * time.Second)
	filter := bson.M{"user_name": username}
	err := ur.collection.FindOne(context, filter).Decode(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}


func (ur *UserRepository) RegisterUser(user domain.User) error {
	context, _ := context.WithTimeout(context.Background(), time.Duration(ur.config.ContextTimeout) * time.Second)
	user.ID = primitive.NewObjectID()
	_, err := ur.collection.InsertOne(context, user)
	if err != nil {
		return err
	}
	return nil
}
func (ur *UserRepository) UpdateUser(user domain.User) error {
	context, _ := context.WithTimeout(context.Background(), time.Duration(ur.config.ContextTimeout) * time.Second)
	filter := bson.M{"email": user.Email}
	update := bson.M{"$set": user}
	_, err := ur.collection.UpdateOne(context, filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (ur *UserRepository) FindUserByID(id string)(domain.User, error){
	var user domain.User
	context, _ := context.WithTimeout(context.Background(), time.Duration(ur.config.ContextTimeout) * time.Second)
	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": objectId}
	err := ur.collection.FindOne(context, filter).Decode(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}