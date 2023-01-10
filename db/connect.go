package db

import (
	"fmt"
	"go-mono/configs"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	dsn := configs.GET("DSN")
	fmt.Println(dsn, "no dsn file")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return db
}
