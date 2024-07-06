package models

type Guild struct {
	ID   string `gorm:"primaryKey"`
	Logs *string
	Mod  *string
}
