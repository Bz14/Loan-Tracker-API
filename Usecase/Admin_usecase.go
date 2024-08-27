package usecases

import (
	"errors"
	domain "loan-tracker/Domain"
	infrastructure "loan-tracker/Infrastructure"
	utils "loan-tracker/Utils"
)

type AdminUseCase struct {
	AdminRepo domain.AdminRepositoryInterface
	UserRepo domain.UserRepositoryInterface
	PassService infrastructure.PasswordService
	Config *infrastructure.Config
}


func NewAdminUseCase(adminRepo domain.AdminRepositoryInterface, passwordService infrastructure.PasswordService, config *infrastructure.Config, userRepo domain.UserRepositoryInterface) *AdminUseCase {
	return &AdminUseCase{
		AdminRepo: adminRepo,
		UserRepo: userRepo,
		PassService: passwordService,
		Config: config,
	}
}


func (ac *AdminUseCase) GetAllUsers(pageNo, pageSize string, user_id string) ([]domain.User, error){
	pageS, pageN, err := utils.PagePaginationValidator(pageSize, pageNo)
	if err != nil{
		return nil, err
	}
	user, err := ac.UserRepo.FindUserByID(user_id)
	if err != nil{
		return nil, err
	}
	if user.Role != "admin"{
		return nil, errors.New("unauthorized: Only admin can access this resource")
	}
	users, err := ac.AdminRepo.GetAllUsers(pageN, pageS)
	if err != nil{
		return nil, errors.New("error getting users")
	}
	return users, nil
}


func (ac *AdminUseCase) DeleteUser(id string, user_id string) (bool, error){
	_, err := ac.UserRepo.FindUserByID(id)
	if err != nil{
		return false, err
	}
	admin, err := ac.UserRepo.FindUserByID(user_id)
	if err != nil || admin.Role != "admin"{
		return false, errors.New("unauthorized: Only admin can access this resource")
	}
	
	err = ac.AdminRepo.DeleteUser(user_id)
	if err != nil{
		return false, err
	}
	return true, nil
}