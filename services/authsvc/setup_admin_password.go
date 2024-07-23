package authsvc

import (
	"context"
	"errors"
	"log"

	"canarails.dev/database/models"
	"canarails.dev/query"
	"canarails.dev/services/authsvc/loginsvc"
	"gorm.io/gorm"
)

// 根据环境变量初始化 Admin 用户的密码
func SetupAdminPassword(ctx context.Context, password string) *models.User {
	salt := loginsvc.CreateRandSalt()

	passwordHash := loginsvc.CreatePasswordHash(password, salt)

	admin, err := query.User.
		WithContext(ctx).
		Where(query.User.Username.Eq("admin")).
		Take()
	if errors.Is(err, gorm.ErrRecordNotFound) {
		admin = &models.User{
			Username:     "admin",
			PasswordHash: passwordHash,
			PasswordSalt: salt,
		}
	}

	admin.PasswordHash = passwordHash
	admin.PasswordSalt = salt
	if err := query.User.WithContext(ctx).Save(admin); err != nil {
		log.Fatalln("初始化 admin 密码失败：", err)
	}

	return admin
}
