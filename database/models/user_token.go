package models

import (
	"time"

	"gorm.io/gorm"
)

type UserToken struct {
	gorm.Model

	Title       string
	Description string
	LastUsedAt  time.Time
	ExpiredAt   time.Time

	User   User
	UserID uint
}
