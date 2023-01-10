package main

import (
	"go-mono/db"

	"gorm.io/gen"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "go-mono",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	db := db.ConnectDB()
	g.UseDB(db)

	g.GenerateAllTable()
	g.Execute()
}
