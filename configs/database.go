package configs

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Databases struct {
	dsn string
	db  *gorm.DB
	err error
}

var DB *gorm.DB

func ConnectDB(dbtype string) {
	con := &Databases{}

	switch dbtype {
	case "mysql":
		con.dsn = GET("DSN")
		fmt.Println(con.dsn, "log dsn")
		con.db, con.err = gorm.Open(mysql.Open(con.dsn), &gorm.Config{})
	default:
		panic("Unknow name database")
	}

	if con.err != nil {
		panic(con.err)
	}

	DB = con.db
	fmt.Println("databases connected")
}
