package controllers

import (
	domain "loan-tracker/Domain"
	"strings"

	"github.com/gin-gonic/gin"
)

type AdminControllers struct{
	AdminUseCase domain.AdminUseCaseInterface
	LoanUseCase domain.LoanUseCaseInterface
}

func NewAdminControllers(adminUseCase domain.AdminUseCaseInterface, loanUsecase domain.LoanUseCaseInterface)*AdminControllers {
	return &AdminControllers{
		AdminUseCase: adminUseCase,
		LoanUseCase: loanUsecase,
	}
}

func (ac *AdminControllers) GetAllUsers(c *gin.Context){
	pageNo := c.Query("pageNo")
	pageSize := c.Query("pageSize")

	if pageNo == ""{
		pageNo = "1"
	}
	if pageSize == ""{
		pageSize = "10"
	}
	user_id := c.GetString("user_id")
	if user_id == "" {
		c.JSON(500, domain.ErrorResponse{
			Message: "Unauthorized: Authorization header required",
			Status:  500,
		})
	}
	users, err := ac.AdminUseCase.GetAllUsers(pageNo, pageSize, user_id)
	if err != nil{
		c.JSON(400, domain.ErrorResponse{
			Message: err.Error(),	
			Status: 400,
		})
		return
	}
	c.JSON(200, domain.SuccessResponse{
		Message: "success",
		Data: users,
		Status: 200,
	})
}


func (ac *AdminControllers) DeleteUser(c *gin.Context){
	id := c.Param("id")
	if id == ""{
		c.JSON(400, domain.ErrorResponse{
			Message: "id is required",
			Status: 400,
		})
		return
	}
	user_id := c.GetString("user_id")
	if user_id == "" {
		c.JSON(500, domain.ErrorResponse{
			Message: "Unauthorized: Authorization header required",
			Status:  500,
		})
	}
	deleted, err := ac.AdminUseCase.DeleteUser(id, user_id)
	if err != nil{
		c.JSON(400, domain.ErrorResponse{
			Message: err.Error(),
			Status: 400,
		})
		return
	}
	c.JSON(200, domain.SuccessResponse{
		Message: "User deleted successfully",
		Data: deleted,
		Status: 200,
	})
}


func (ac *AdminControllers) GetAllLoans(c *gin.Context){
	status := c.Param("status")
	if status == ""{
		status = ""
	}
	order := c.Param("order")
	if order == ""{
		if strings.ToLower(status) == "reviewed"{
			order = "desc"
		} else if strings.ToLower(status) == "pending"{
			order = "asc"
		}
	}
	user_id := c.GetString("user_id")
	if user_id == "" {
		c.JSON(500, domain.ErrorResponse{
			Message: "Unauthorized: Authorization header required",
			Status:  500,
		})
	}
	loans, err := ac.LoanUseCase.GetAllLoans(status, order, user_id)
	if err != nil{
		c.JSON(400, domain.ErrorResponse{
			Message: err.Error(),	
			Status: 400,
		})
		return
	}
	c.JSON(200, domain.SuccessResponse{
		Message: "Available Loans",
		Data: loans,
		Status: 200,
	})
}

