package models

import "gorm.io/gorm"

type App struct {
	gorm.Model

	Title       string
	Description string
	Hostnames   []string `gorm:"serializer:json"`

	AppVariants []AppVariant
}
