package repository

import (
	"rsvp-system/internal/models"

	"gorm.io/gorm"
)

type EventRepository interface {
	Create(event *models.Event) error
	GetAll(userID uint) ([]models.Event, error)
	GetByID(id uint) (*models.Event, error)
	Delete(id uint) error
}

type eventRepo struct {
	db *gorm.DB
}

func NewEventRepository(db *gorm.DB) EventRepository {
	return &eventRepo{db}
}

func (r *eventRepo) Create(event *models.Event) error {
	return r.db.Create(event).Error
}

func (r *eventRepo) GetAll(userID uint) ([]models.Event, error) {
	var events []models.Event
	err := r.db.Where("user_id = ?", userID).Find(&events).Error
	return events, err
}

func (r *eventRepo) GetByID(id uint) (*models.Event, error) {
	var event models.Event
	err := r.db.First(&event, id).Error
	return &event, err
}

func (r *eventRepo) Delete(id uint) error {
	return r.db.Delete(&models.Event{}, id).Error
}
