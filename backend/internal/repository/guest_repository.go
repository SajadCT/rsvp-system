package repository

import (
	"rsvp-system/internal/models"

	"gorm.io/gorm"
)

type GuestRepository interface {
	Create(guest *models.Guest) error
	GetByEventID(eventID uint) ([]models.Guest, error)
	GetByID(id uint) (*models.Guest, error)
	UpdateStatus(guest *models.Guest, status string) error
	GetGuestWithEvent(id uint) (*models.Guest, error)
	CountByStatus(eventID uint, status string) (int64, error)
}

type guestRepo struct {
	db *gorm.DB
}

func NewGuestRepository(db *gorm.DB) GuestRepository {
	return &guestRepo{db}
}

func (r *guestRepo) Create(guest *models.Guest) error {
	return r.db.Create(guest).Error
}

func (r *guestRepo) GetByEventID(eventID uint) ([]models.Guest, error) {
	var guests []models.Guest
	err := r.db.Where("event_id = ?", eventID).Find(&guests).Error
	return guests, err
}

func (r *guestRepo) GetByID(id uint) (*models.Guest, error) {
	var guest models.Guest
	err := r.db.First(&guest, id).Error
	return &guest, err
}

func (r *guestRepo) UpdateStatus(guest *models.Guest, status string) error {
	guest.Status = status
	return r.db.Save(guest).Error
}

func (r *guestRepo) GetGuestWithEvent(id uint) (*models.Guest, error) {
	var guest models.Guest
	err := r.db.Preload("Event").First(&guest, id).Error
	return &guest, err
}

func (r *guestRepo) CountByStatus(eventID uint, status string) (int64, error) {
	var count int64
	query := r.db.Model(&models.Guest{}).Where("event_id = ?", eventID)
	if status != "" {
		query = query.Where("status = ?", status)
	}
	err := query.Count(&count).Error
	return count, err
}
