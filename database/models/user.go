package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Username     string
	PasswordHash []byte
	PasswordSalt []byte

	UserTokens []UserToken
}
