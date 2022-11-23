package config

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func ConnectDatabase() *gorm.DB {
	DBMS := "postgres"
	HOST := "go-trial-db"
	PORT := "5432"
	USER := "user"
	DBNAME := "pg-db"
	PASS := "pass"
	CONNECTION := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", HOST, PORT, USER, DBNAME, PASS)
	db, err := gorm.Open(DBMS, CONNECTION)
	if err != nil {
		// 1秒ごとに10回繰り返す
		for count := 0; count <= 10; count++ {
			if err == nil {
				break
			}
			fmt.Println(".")
			time.Sleep(time.Second)
			db, err = gorm.Open(DBMS, CONNECTION)
		}

		if err != nil {
			fmt.Println("db connection failed")
			panic(err)
		}
	}
	fmt.Println("db connection success")
	return db
}