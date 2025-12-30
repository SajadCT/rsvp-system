package models

import "gorm.io/gorm"

type Event struct {
	gorm.Model
	Title       string `gorm:"not null"`
	Description string
	Date        string `gorm:"not null"`
	Location    string `gorm:"not null"`

	UserID uint `gorm:"not null"`
	User   User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	Guests []Guest
}
