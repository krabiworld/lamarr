package models

import "gorm.io/gorm"

type Warn struct {
	gorm.Model
	GuildID  string
	MemberID string
	Reason   string
}
