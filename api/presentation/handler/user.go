package handler

import (
	"go-trial/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userUsecase usecase.UserUsecase
}

func SetUserHandler(uu usecase.UserUsecase) UserHandler {
	return UserHandler{userUsecase: uu}
}

func (uh UserHandler) GetList(c *gin.Context) {
	users, err := uh.userUsecase.GetList()
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}