package models

import "gorm.io/gorm"

type Stats struct {
	gorm.Model
	ExecutedCommands uint64
}
