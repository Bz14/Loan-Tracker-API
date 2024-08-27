package usecases

import (
	"errors"
	domain "loan-tracker/Domain"
	infrastructure "loan-tracker/Infrastructure"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type UserUseCase struct {
	UserRepo domain.UserRepositoryInterface
	PassService infrastructure.PasswordService
	Config *infrastructure.Config
}


func NewUserUseCase(userRepo domain.UserRepositoryInterface, passwordService infrastructure.PasswordService, config *infrastructure.Config) *UserUseCase {
	return &UserUseCase{
		UserRepo: userRepo,
		PassService: passwordService,
		Config: config,
	}
}

func (uc *UserUseCase) RegisterUser(user domain.User) error {
	if user.Email == "" || user.User_Name == "" || user.Password == "" {
		return errors.New("all fields are required")
	}
	if infrastructure.ValidateEmail(user.Email) != nil {
		return errors.New("invalid email format")
	}
	if infrastructure.ValidatePassword(user.Password) != nil {
		return errors.New("invalid password format")
	}

	existingUser, err := uc.UserRepo.FindUserByEmail(user.Email)

	if err == nil && existingUser.Email != "" && !existingUser.IsVerified{
		return errors.New("user already exists: Verify your account")
	}else if err == nil{
		return errors.New("user already exists: Try with another email")
	}

	hashedPassword, _ := uc.PassService.HashPassword(user.Password)
	user.Password = hashedPassword

	token, err  := infrastructure.GenerateVerificationToken()
	if err != nil{
		return errors.New("error generating verification token")
	}
	err = infrastructure.SendVerificationEmail(user.Email, token)
	if err != nil {
		return errors.New("error sending verification email")
	}

	user.IsVerified = false
	user.Created_At = time.Now()
	user.VerificationToken = token
	user.VerificationExpires = time.Now().Add(time.Hour * 24)
	user.Role = "user"
	err = uc.UserRepo.RegisterUser(user)
	if err != nil {
		return errors.New("error creating user")
	}
	return nil
}



func (uc *UserUseCase) VerifyEmail(email string, token string)error{
	user, err := uc.UserRepo.FindUserByEmail(email)
	if err != nil {
		return errors.New("user not found")
	}

	if user.VerificationToken != token {
		return errors.New("invalid verification token")
	}
	if user.VerificationExpires.Before(time.Now()) {
		return errors.New("verification token expired. Please request a new one")
	}

	user.IsVerified = true
	user.VerificationToken = ""
	user.VerificationExpires = time.Time{}
	err = uc.UserRepo.UpdateUser(user)
	if err != nil {
		return errors.New("error verifying user")
	}
	return nil
}

func (uc *UserUseCase) Login(user domain.User)(domain.LoginResponse, error){
	var newUser domain.User
	var err error
	if user.Email == "" || user.Password == "" {
		return domain.LoginResponse{}, errors.New("all fields are required")
	}
	if user.Email != ""{
		newUser, err = uc.UserRepo.FindUserByEmail(user.Email)	
	}else if user.User_Name != ""{
		newUser, err = uc.UserRepo.FindUserByUserName(user.User_Name)
	}
	if err != nil{
		return domain.LoginResponse{}, errors.New("user not found")
	}
	if !newUser.IsVerified{
		return domain.LoginResponse{}, errors.New("user not verified. Please Verify Your Account")
	}
	if !uc.PassService.ComparePassword(user.Password, newUser.Password){
		return domain.LoginResponse{}, errors.New("invalid credentials: Password does not match")
	}
	accessToken, err := uc.CreateAccessToken(&newUser, uc.Config.AccessTokenSecret, uc.Config.AccessTokenExpiryHour)
	if err != nil {
		return domain.LoginResponse{}, errors.New("error creating access token")
	}
	refreshToken, err  := uc.CreateRefreshToken(&newUser, uc.Config.RefreshTokenSecret, uc.Config.RefreshTokenExpiryHour)
	if err != nil {
		return domain.LoginResponse{}, errors.New("error creating refresh token")
	}
	return domain.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}


func (uc *UserUseCase) CreateAccessToken(user *domain.User, secret string, expiry int) (accessToken string, err error) {
	exp := time.Now().Add(time.Hour * time.Duration(expiry))
	claims := &domain.JwtCustomClaims{
		ID: user.ID.Hex(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}
	return infrastructure.CreateToken(claims, secret)
}

func (uc *UserUseCase) CreateRefreshToken(user *domain.User, secret string, expiry int) (refreshToken string, err error) {
	exp := time.Now().Add(time.Hour * time.Duration(expiry))

	claims := &domain.JwtCustomClaims{
		ID: user.ID.Hex(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	}
	return infrastructure.CreateToken(claims, secret)
}


func (uc *UserUseCase) RefreshToken(request domain.RefreshTokenRequest, user_id string) (domain.RefreshTokenResponse, error) {
	id, err := infrastructure.ExtractIDFromToken(request.RefreshToken, uc.Config.RefreshTokenSecret)
	if err != nil {
		return domain.RefreshTokenResponse{}, errors.New("user not found")
	}
	valid, err := infrastructure.IsAuthorized(request.RefreshToken, uc.Config.RefreshTokenSecret)
	if err != nil {
		return domain.RefreshTokenResponse{}, errors.New("user not found")
	}

	if id != user_id && !valid {
		return domain.RefreshTokenResponse{}, errors.New("session expired")
	}

	user, err := uc.UserRepo.FindUserByID(user_id)
	if err != nil {
		return domain.RefreshTokenResponse{}, errors.New("user not found")
	}
	accessToken, err := uc.CreateAccessToken(&user, uc.Config.AccessTokenSecret, uc.Config.AccessTokenExpiryHour)
	if err != nil {
		return domain.RefreshTokenResponse{}, errors.New(err.Error())
	}
	return domain.RefreshTokenResponse{
		AccessToken:  accessToken,
	}, nil
}


func (uc *UserUseCase) GetUserProfile(id string)(domain.UserProfile, error){
	user, err := uc.UserRepo.FindUserByID(id)
	if err != nil {
		return domain.UserProfile{}, errors.New("user not found")
	}
	return domain.UserProfile{
		ID: user.ID,
		User_Name: user.User_Name,
		Email: user.Email,
		Contact: user.Contact,
		Created_At: user.Created_At,
	}, nil
}

func (uc *UserUseCase) ResetPassword(email string, user_id string) error{
	user, err := uc.UserRepo.FindUserByEmail(email)
	if err != nil {
		return errors.New("user not found")
	}
	if user.ID.Hex() != user_id {
		return errors.New("unauthorized: User not found")
	}

	token, err := infrastructure.GenerateVerificationToken()
	if err != nil {
		return errors.New("error generating token")
	}
	err = infrastructure.SendResetPasswordVerificationEmail(email, token)
	if err != nil {
		return errors.New("error sending email")
	}

	user.ResetPasswordToken = token
	user.ResetPasswordExpires = time.Now().Add(time.Hour * 24)
	err = uc.UserRepo.UpdateUser(user)
	if err != nil {
		return errors.New("error updating user")
	}
	return nil
}


func (uc *UserUseCase) ResetPasswordVerify(email string, token string, user_id string, password string) error{
	user, err := uc.UserRepo.FindUserByEmail(email)
	if err != nil {
		return errors.New("user not found")
	}
	if user.ID.Hex() != user_id {
		return errors.New("unauthorized: User not found")
	}

	if err := infrastructure.ValidatePassword(password); err != nil {
		return errors.New("invalid password format")
	}

	if user.ResetPasswordToken != token {
		return errors.New("invalid token")
	}
	if user.ResetPasswordExpires.Before(time.Now()) {
		return errors.New("token expired")
	}
	newPassword, err := uc.PassService.HashPassword(password)
	if err != nil{
		return errors.New("error hashing password")
	}
	user.Password = newPassword
	user.ResetPasswordToken = ""
	user.ResetPasswordExpires = time.Time{}

	err = uc.UserRepo.UpdateUser(user)
	return nil
}