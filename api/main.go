package main

import (
	"fmt"

	// "github.com/jinzhu/gorm"
	// "github.com/jinzhu/gorm"
	// "go-trial/infra/config"
	// "github.com/gin-gonic/gin"
	// "net/http"
)

////////////////////////////
type Human struct {
	name string
}

type Actor interface {
	// actor.Walk() を呼び出し可能にする
	Walk()
}

type Tester struct {
	// 実際に代入されるのは user
	actor Actor
}

func (tester Tester) Walk() {
	// tester.actor には human が代入されている
	// actor.Walk() は func (human Human) Walk() で処理される
	tester.actor.Walk()
}

func (human Human) Walk() {
	fmt.Println(human.name + " is walking")
}

func test() {
	human := Human{name: "Taro"}
	// actor に human を代入する
	tester := Tester{actor: human}
	tester.Walk()
}

/////////////////////////////////////////
// handler 層
type Handler struct {
	usecase Usecase
}

func (h Handler) GetList() []User {
	users := h.usecase.GetList()
	return users
}


/////////////////////////////////////////
// usecase 層
type Usecase struct {
	// 実際に挿入されるのは dbRepo
	repo Repo
}

func (usecase Usecase) GetList() []User {
	// usecase.repo には dbRepo が代入されている
	// repo.GetList() は infra 層の func (dbRepo DbRepo) GetList() で処理される
	users := usecase.repo.GetList()
	return users
}

/////////////////////////////////////////
// domain 層
type User struct {
	name string
}
type Repo interface {
	// repo.GetList() を呼び出し可能にする
	GetList() []User
}

/////////////////////////////////////////
// infra 層
type DbRepo struct {
	dbType string
}

func (dbRepo DbRepo) GetList() []User {
	fmt.Println("user is getting list from " + dbRepo.dbType)

	users := []User{{name: "ichiro"}, {name: "Jiro"}}
	return users
}

func main() {
	// db := config.ConnectDatabase()
	// defer db.Close()

	// router := gin.Default()
	// router.GET("/", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "Initial API",
	// 	})
	// })
	// router.Run(":3000")
	dbRepo := DbRepo{dbType: "postgres"}
	// repo に dbRepo を代入することで repo に渡ったDB処理を infra 層で行う
	usecase := Usecase{repo: dbRepo}
	handler := Handler{usecase: usecase}
	users := handler.GetList()
	fmt.Println(users)
}