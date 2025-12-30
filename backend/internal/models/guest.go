package models

import "gorm.io/gorm"

type Guest struct {
	gorm.Model

	Name string `gorm:"not null"`

	Email string `gorm:"not null;uniqueIndex:idx_event_email"`

	Status string `gorm:"default:'Pending';not null"`

	EventID uint `gorm:"not null;uniqueIndex:idx_event_email"`

	Event Event `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
