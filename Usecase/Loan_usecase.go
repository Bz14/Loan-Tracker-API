package usecases

import (
	"errors"
	domain "loan-tracker/Domain"
	infrastructure "loan-tracker/Infrastructure"
	"time"
)

type LoanUseCase struct {
	LoanRepo domain.LoanRepositoryInterface
	UserRepo domain.UserRepositoryInterface
	PassService infrastructure.PasswordService
	Config *infrastructure.Config
}


func NewLoanUseCase(loanRepo domain.LoanRepositoryInterface, passwordService infrastructure.PasswordService, config *infrastructure.Config, userRepo domain.UserRepositoryInterface) *LoanUseCase {
	return &LoanUseCase{
		LoanRepo: loanRepo,
		UserRepo: userRepo,
		PassService: passwordService,
		Config: config,
	}
}


func (lu *LoanUseCase) CreateLoan(loan domain.Loan, user_id string) error {
	loan.LoanStatus = "pending"
	loan.Created_at = time.Now()
	if loan.UserId.Hex() != user_id {
		return errors.New("Unauthorized")
	}
	if loan.Amount < 1 {
		return errors.New("Invalid amount")
	}
	err := lu.LoanRepo.CreateLoan(loan)
	if err != nil {
		return err
	}
	return nil
}


func (lu *LoanUseCase) CheckLoanStatus(id string, user_id string) (string, error){
	loan, err := lu.LoanRepo.FindLoanByID(id)
	if err != nil{
		return "", errors.New("loan Not found")
	}
	if loan.UserId.Hex() != user_id{
		return "", errors.New("unauthorized: User can access this loan")
	}
	return loan.LoanStatus, nil
}


func (lu *LoanUseCase) GetAllLoans(status string, order string, user_id string) ([]domain.Loan, error){
	user, err := lu.UserRepo.FindUserByID(user_id)
	if err != nil{
		return nil, errors.New("user not found")
	}
	if user.Role != "admin"{
		return nil, errors.New("unauthorized: Only admin can access this resource")
	}
	loans, err := lu.LoanRepo.GetAllLoans(status, order)
	if err != nil{
		return nil, errors.New("can not retrieve loans")
	}
	
	return loans, nil
}