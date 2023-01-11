package configs

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Databases struct {
	dsn string
	db  *gorm.DB
	err error
}

func ConnectDB(dbtype string) *gorm.DB {
	con := &Databases{}

	switch dbtype {
	case "mysql":
		con.dsn = GET("DSN")
		con.db, con.err = gorm.Open(mysql.Open(con.dsn), &gorm.Config{})
	default:
		panic("Unknow name database")
	}

	if con.err != nil {
		panic(con.err)
	}

	return con.db
}
