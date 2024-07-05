package models

type Guild struct {
	ID     string `gorm:"primaryKey"`
	Prefix string `gorm:"not null;default:!"`
	Logs   *string
	Mod    *string
}
