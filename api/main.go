package main

import (
	"go-trial/infra/config"
	"go-trial/infra/repoimple"
	"go-trial/presentation/handler"
	"go-trial/usecase"
	"github.com/gin-gonic/gin"
)

func main() {
	// db
	db := config.ConnectDatabase()
	// db := config.MigrateDatabase()
	defer db.Close()

	// routing
	router := gin.Default()

	// user
	userRepoimple := repoimple.SetUserRepoimple(db)
	userUsecase := usecase.SetUserUsecase(userRepoimple)
	userHandler := handler.SetUserHandler(userUsecase)
	userRouter := router.Group("/user")
	{
		userRouter.GET("/", userHandler.GetList)
		userRouter.POST("/", userHandler.Create)
		userRouter.GET("/:id", userHandler.GetSingle)
	}

	router.Run(":3000")
}