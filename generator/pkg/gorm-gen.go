package generator

import (
	"go-mono/configs"

	"gorm.io/gen"
)

func GormGen() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "go-mono",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	db := configs.ConnectDB("mysql")
	g.UseDB(db)

	g.GenerateAllTable()
	g.Execute()
}
