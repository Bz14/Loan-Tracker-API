package controllers

import (
	"fmt"
	domain "loan-tracker/Domain"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type LoanControllers struct{
	LoanUseCase domain.LoanUseCaseInterface
}

func NewLoanControllers(loanUseCase domain.LoanUseCaseInterface) *LoanControllers {
	return &LoanControllers{
		LoanUseCase: loanUseCase,
	}
}


func (lc *LoanControllers) CreateLoan(c *gin.Context) {
	var loan domain.Loan

	err := c.BindJSON(&loan)
	if err != nil {
		c.JSON(400, domain.ErrorResponse{
			Message: "Invalid request",
			Status:  400,
		})
		return
	}
	validate := validator.New()

	if err := validate.Struct(loan); err != nil {
		c.JSON(400, domain.ErrorResponse{
			Message: "Invalid request",
			Status:  400,
		})
		return
	}
	user_id := c.GetString("user_id")
	if user_id == "" {
		c.JSON(500, domain.ErrorResponse{
			Message: "Unauthorized",
			Status:  500,
		})
		return
	}

	err = lc.LoanUseCase.CreateLoan(loan, user_id)
	if err != nil{
		fmt.Println(err.Error())
		c.JSON(400, domain.ErrorResponse{
			Message: err.Error(),
			Status:  400,
		})
		return
	}

	c.JSON(200, domain.SuccessResponse{
		Message: "Loan application success",
		Status:  200,
	})
}


func (lc *LoanControllers) CheckLoanStatus(c *gin.Context){
	id := c.Param("id")
	if id == ""{
		c.JSON(500, domain.ErrorResponse{
			Message: "Loan Id required",
			Status:  500,
		})
		return
	}
	user_id := c.GetString("user_id")
	if user_id == "" {
		c.JSON(500, domain.ErrorResponse{
			Message: "Unauthorized",
			Status:  500,
		})
		return
	}

	status, err := lc.LoanUseCase.CheckLoanStatus(id, user_id)
	if err != nil{
		c.JSON(400, domain.ErrorResponse{
			Message: err.Error(),
			Status:  400,
		})
		return
	}

	c.JSON(200, domain.SuccessResponse{
		Message: "Loan status detail",
		Data : map[string]string{
			"Loan ID" : id,
			"Status": status,
		},
		Status:  200,
	})
	
}