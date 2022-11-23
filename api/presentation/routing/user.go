package routing

import (
	"go-trial/infra/repoimple"
	"go-trial/presentation/handler"
	"go-trial/usecase"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func UserRouting(db *gorm.DB, router *gin.Engine) {
	// DI: DependencyInjection: 依存性の注入
	userRepoimple := repoimple.SetUserRepoimple(db)
	// 変数の userRepo に userRepoimple を注入する
	userUsecase := usecase.SetUserUsecase(userRepoimple)
	userHandler := handler.SetUserHandler(userUsecase)
	userRouter := router.Group("/user")
	{
		userRouter.GET("/", userHandler.GetList)
		userRouter.POST("/", userHandler.Create)
		userRouter.GET("/:id", userHandler.GetSingle)
		userRouter.PUT("/:id", userHandler.Update)
		userRouter.DELETE("/:id", userHandler.Delete)
	}
}