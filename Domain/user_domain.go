package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID                   primitive.ObjectID   `bson:"_id,omitempity" json:"id" `
	User_Name            string               `bson:"user_name"  json:"user_name"`
	Email 				 string 			  `bson:"email" validate:"required,email" json:"email"`
    Password             string               `bson:"password" json:"password" validate:"required"`
    Contact              string               `bson:"contact" json:"contact"`
	IsVerified			 bool 				  `bson:"is_verified" json:"is_verified"`
	Created_At		     time.Time			  `bson:"created_at" json:"created_at"`
    ResetPasswordToken   string               `bson:"reset_password_token" json:"reset_password_token"`
    ResetPasswordExpires time.Time            `bson:"reset_password_expires" json:"reset_password_expires"`
	VerificationToken    string               `bson:"verification_token" json:"verification_token"`
	VerificationExpires  time.Time            `bson:"verification_expires" json:"verification_expires"`
	Role 			   	 string               `bson:"role" json:"role"`
}


type UserUseCaseInterface interface {
	RegisterUser(user User) error
	VerifyEmail(email string, token string) error
	Login(user User)(LoginResponse, error)
	CreateAccessToken(user *User, secret string, expiry int) (accessToken string, err error)
	CreateRefreshToken(user *User, secret string, expiry int) (refreshToken string, err error)
	RefreshToken(request RefreshTokenRequest, user_id string) (RefreshTokenResponse, error)
	GetUserProfile(id string)(UserProfile, error)
	ResetPassword(email string, user_id string)error
	ResetPasswordVerify(email string, token string, user_id string, password string) error
}


type UserRepositoryInterface interface {
	RegisterUser(user User) error
	FindUserByEmail(email string) (User, error)
	UpdateUser(user User) error
	FindUserByUserName(username string) (User, error)
	FindUserByID(id string)(User, error)
}


type AdminUseCaseInterface interface {
	GetAllUsers(pageNo, pageSize string, user_id string) ([]User, error)
	DeleteUser(id string, user_id string) (bool, error)
}

type AdminRepositoryInterface interface {
	GetAllUsers(pageNo, pageSize int64) ([]User, error)
	DeleteUser(id string) error
}