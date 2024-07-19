package main

import (
	"log"

	"canarails.dev/apis"
	"canarails.dev/apis/genapi"
	"canarails.dev/database"
	"github.com/labstack/echo/v4"
)

func main() {
	database.GetDb()

	app := echo.New()

	ssi := apis.New()
	si := genapi.NewStrictHandler(ssi, nil)

	genapi.RegisterHandlers(app, si)

	log.Fatal(app.Start(":3000"))
}
