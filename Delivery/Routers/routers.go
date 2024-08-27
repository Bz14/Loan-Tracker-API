package routers

import (
	controllers "loan-tracker/Delivery/Controllers"
	infrastructure "loan-tracker/Infrastructure"
	repository "loan-tracker/Repository"
	useCase "loan-tracker/Usecase"

	"github.com/gin-gonic/gin"
)


func Routers(server *gin.Engine, db *infrastructure.Db, config *infrastructure.Config) {
	user_collection := db.CreateDb(config.DatabaseUrl, config.DbName, config.UserCollection)
	loan_collection := db.CreateDb(config.DatabaseUrl, config.DbName, config.LoanCollection)

	user_repository := repository.NewUserRepository(user_collection, config)
	loan_repository := repository.NewLoanRepository(loan_collection, config)
	admin_repository := repository.NewAdminRepository(user_collection, config)

	password_service := infrastructure.NewPasswordService()
	user_useCase := useCase.NewUserUseCase(user_repository, *password_service, config)
	loan_usecase := useCase.NewLoanUseCase(loan_repository, *password_service, config, user_repository)
	admin_useCase := useCase.NewAdminUseCase(admin_repository, *password_service, config, user_repository)

	userControllers := controllers.NewUserControllers(user_useCase)

	adminControllers := controllers.NewAdminControllers(admin_useCase)

	loan_controller := controllers.NewLoanControllers(loan_usecase)
	
	authMiddleWare := infrastructure.NewAuthMiddleware(*config).AuthenticationMiddleware()


	nonAuth := server.Group("users")
	nonAuth.POST("/register", userControllers.RegisterUser)
	nonAuth.GET("/verify-email", userControllers.VerifyEmail)
	nonAuth.POST("/login", userControllers.Login)


	adminRoute := server.Group("admin")
	adminRoute.GET("/users", authMiddleWare, adminControllers.GetAllUsers)
	adminRoute.DELETE("/users/:id", authMiddleWare, adminControllers.DeleteUser)
	
	
	auth := server.Group("users")
	auth.GET("/profile", authMiddleWare, userControllers.GetUserProfile)
	auth.POST("/password-reset", authMiddleWare, userControllers.ResetPassword)
	auth.POST("/password-update", authMiddleWare, userControllers.ResetPasswordVerify)

	loanRoute := server.Group("loans")
	loanRoute.POST("", authMiddleWare, loan_controller.CreateLoan)


	tokenGroup := server.Group("token")
	tokenGroup.POST("/refresh", authMiddleWare, userControllers.RefreshToken)

}
