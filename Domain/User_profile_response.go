package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserProfile struct {
	ID         primitive.ObjectID `bson:"_id,omitempity" json:"id" `
	User_Name  string             `bson:"user_name"  json:"user_name"`
	Email      string             `bson:"email" validate:"required,email" json:"email"`
	Contact    string             `bson:"contact" json:"contact"`
	Created_At time.Time          `bson:"created_at" json:"created_at"`
}