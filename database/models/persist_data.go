package models

import "gorm.io/gorm"

type PersistData struct {
	gorm.Model

	Key   string
	Value string
}
