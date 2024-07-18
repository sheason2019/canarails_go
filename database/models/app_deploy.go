package models

import "gorm.io/gorm"

type AppDeploy struct {
	gorm.Model

	Replicas  int
	ImageName string

	AppVariant   AppVariant
	AppVariantID uint
}
