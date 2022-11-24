package handler

import (
	"go-trial/lib/errorhandle"
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
		parsedErr := errorhandle.InitParsedError(err)
		c.JSON(parsedErr.Status, gin.H{
			"error": parsedErr.OrgErr,
			"status": parsedErr.Status,
			"message": parsedErr.Message,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

func (uh UserHandler) Create(c *gin.Context) {
	name := c.PostForm("name")
	user, err := uh.userUsecase.Create(name)
	if err != nil {
		parsedErr := errorhandle.InitParsedError(err)
		c.JSON(parsedErr.Status, gin.H{
			"error": parsedErr.OrgErr,
			"status": parsedErr.Status,
			"message": parsedErr.Message,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func (uh UserHandler) GetSingle(c *gin.Context) {
	id := c.Param("id")
	user, err := uh.userUsecase.GetSingle(id)
	if err != nil {
		parsedErr := errorhandle.InitParsedError(err)
		c.JSON(parsedErr.Status, gin.H{
			"error": parsedErr.OrgErr,
			"status": parsedErr.Status,
			"message": parsedErr.Message,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func (uh UserHandler) Update(c *gin.Context) {
	id := c.Param("id")
	name := c.PostForm("name")
	user, err := uh.userUsecase.Update(id, name)
	if err != nil {
		parsedErr := errorhandle.InitParsedError(err)
		c.JSON(parsedErr.Status, gin.H{
			"error": parsedErr.OrgErr,
			"status": parsedErr.Status,
			"message": parsedErr.Message,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func (uh UserHandler) Delete(c *gin.Context) {
	id := c.Param("id")
	user, err := uh.userUsecase.Delete(id)
	if err != nil {
		parsedErr := errorhandle.InitParsedError(err)
		c.JSON(parsedErr.Status, gin.H{
			"error": parsedErr.OrgErr,
			"status": parsedErr.Status,
			"message": parsedErr.Message,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}
