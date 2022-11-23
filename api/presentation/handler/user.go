package handler

import (
	"fmt"
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
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

func (uh UserHandler) Create(c *gin.Context) {
	name := c.PostForm("name")
	user, err := uh.userUsecase.Create(name)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func (uh UserHandler) GetSingle(c *gin.Context) {
	id := c.Param("id")
	user, err := uh.userUsecase.GetSingle(id)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func (uh UserHandler) Update(c *gin.Context) {
	id := c.Param("id")
	name := c.PostForm("name")
	fmt.Println("---------------------------------")
	fmt.Println(id)
	fmt.Println(name)
	fmt.Println("---------------------------------")
	user, err := uh.userUsecase.Update(id, name)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"error": err.Error(),
		})
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}
