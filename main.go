package main

import (
	"context"
	"log"
	"os"

	"canarails.dev/apis"
	"canarails.dev/apis/genapi"
	"canarails.dev/database"
	"canarails.dev/query"
	"canarails.dev/services/authsvc"
	"canarails.dev/services/envsvc"
	"github.com/labstack/echo/v4"
)

func main() {
	// 确认环境变量已初始化
	envsvc.EnsureEnvironmentInitial()
	// 确认数据库连接已建立
	query.SetDefault(database.GetDb())
	// 初始化 admin 用户密码
	authsvc.SetupAdminPassword(context.Background(), os.Getenv(envsvc.ADMIN_PASSWORD))

	// 初始化 Api 服务
	app := echo.New()

	ssi := apis.New()
	si := genapi.NewStrictHandler(ssi, nil)

	genapi.RegisterHandlers(app, si)

	log.Fatal(app.Start(":8080"))
}
