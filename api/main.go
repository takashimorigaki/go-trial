package main

import (
	"go-trial/infra/config"
	"go-trial/presentation/routing"

	"github.com/gin-gonic/gin"
)

func main() {
	// db
	db := config.ConnectDatabase()
	// db := config.MigrateDatabase()
	defer db.Close()

	// routing
	router := gin.Default()
	routing.UserRouting(db, router)
	router.Run(":3000")
}