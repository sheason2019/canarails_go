package main

import (
	"canarails.dev/database"
	"canarails.dev/database/models"
	"gorm.io/gen"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./query",
		Mode:    gen.WithDefaultQuery | gen.WithQueryInterface,
	})

	g.UseDB(database.GetDb())

	g.ApplyBasic(models.Models...)

	g.Execute()
}
