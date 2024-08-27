package usecases

import (
	"errors"
	domain "loan-tracker/Domain"
	infrastructure "loan-tracker/Infrastructure"
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
	if loan.UserId != user_id {
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