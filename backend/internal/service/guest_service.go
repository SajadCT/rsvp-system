package service

import (
	"errors"
	"rsvp-system/internal/models"
	"rsvp-system/internal/repository"
)

type GuestService interface {
	InviteGuest(guest *models.Guest) error
	GetGuestsForEvent(eventID uint) ([]models.Guest, error)
	RSVPGuest(guestID uint, status string) error
	GetGuestDetails(guestID uint) (*models.Guest, error)
}

type guestService struct {
	repo      repository.GuestRepository
	eventRepo repository.EventRepository
}

func NewGuestService(repo repository.GuestRepository, eventRepo repository.EventRepository) GuestService {
	return &guestService{repo, eventRepo}
}

func (s *guestService) InviteGuest(guest *models.Guest) error {

	_, err := s.eventRepo.GetByID(guest.EventID)
	if err != nil {
		return errors.New("event not found")
	}

	guest.Status = "Pending"
	return s.repo.Create(guest)
}

func (s *guestService) GetGuestsForEvent(eventID uint) ([]models.Guest, error) {
	return s.repo.GetByEventID(eventID)
}

func (s *guestService) RSVPGuest(guestID uint, status string) error {
	guest, err := s.repo.GetByID(guestID)
	if err != nil {
		return err
	}
	return s.repo.UpdateStatus(guest, status)
}

func (s *guestService) GetGuestDetails(guestID uint) (*models.Guest, error) {
	return s.repo.GetGuestWithEvent(guestID)
}
