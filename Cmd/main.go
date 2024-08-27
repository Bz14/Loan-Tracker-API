package main

import (
	"fmt"

	routers "loan-tracker/Delivery/Routers"
	infrastructure "loan-tracker/Infrastructure"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	config, err := infrastructure.LoadEnv()
	if err != nil {
		fmt.Print("Error in env.load")
	}
	db := infrastructure.NewDatabase()
	
	routers.Routers(server, db, config)

	server.Run(fmt.Sprintf(":%d", config.Port))

}