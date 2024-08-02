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
	"canarails.dev/services/gatewaysvc"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// 确认环境变量已初始化
	envsvc.EnsureEnvironmentInitial()
	// 确认数据库连接已建立
	query.SetDefault(database.GetDb())
	// 初始化 admin 用户密码
	authsvc.SetupAdminPassword(context.Background(), os.Getenv(envsvc.ADMIN_PASSWORD))
	// 同步 Gateway
	err := gatewaysvc.Reconciliation(context.Background(), query.Q)
	if err != nil {
		log.Fatalln("初始化 Canarails 失败：", err)
	}

	// 初始化 Api 服务
	app := echo.New()

	app.Use(authsvc.AuthMiddleware)

	ssi := apis.New()
	si := genapi.NewStrictHandler(ssi, nil)

	genapi.RegisterHandlers(app, si)

	app.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		HTML5: true,
		Root:  "wwwroot",
		Index: "html/main/index.html",
	}))

	log.Fatal(app.Start(":3000"))
}
