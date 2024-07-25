package models

import "gorm.io/gorm"

type AppVariant struct {
	gorm.Model

	Title       string
	Description string
	ExposePort  int
	Matches     []AppVariantMatch `gorm:"serializer:json"`

	App   App
	AppID uint
}

type AppVariantMatch struct {
	Header string
	Value  string
}
