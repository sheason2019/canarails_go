package envsvc

import (
	"log"
	"os"
	"strings"
)

// 确认程序必须的环境变量已进行初始化
func EnsureEnvironmentInitial() {
	envNames := []string{
		ADMIN_PASSWORD,
		DATABASE_URL,
	}

	larkList := []string{}

	for _, name := range envNames {
		if len(os.Getenv(name)) == 0 {
			larkList = append(larkList, name)
		}
	}

	if len(larkList) != 0 {
		log.Fatalln("程序初始化失败，缺少环境变量:", strings.Join(larkList, ","))
	}
}
