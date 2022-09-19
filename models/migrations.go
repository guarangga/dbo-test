package models

type Migration struct {
	Id        int    `gorm:"primaryKey"`
	Migration string `gorm:"size:255;not null"`
	Batch     int    `gorm:"not null"`
}
