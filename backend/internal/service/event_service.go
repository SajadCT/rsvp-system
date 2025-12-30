package service

import (
	"errors"
	"rsvp-system/internal/models"
	"rsvp-system/internal/repository"
)

type EventService interface {
	CreateEvent(event *models.Event, userID uint) error
	GetAllEvents(userID uint) ([]models.Event, error)

	DeleteEvent(id uint) error
	GetEventStats(eventID uint) (map[string]int64, error)
}

type eventService struct {
	repo      repository.EventRepository
	guestRepo repository.GuestRepository
}

func NewEventService(repo repository.EventRepository, guestRepo repository.GuestRepository) EventService {
	return &eventService{repo, guestRepo}
}

func (s *eventService) CreateEvent(event *models.Event, userID uint) error {
	if event.Title == "" {
		return errors.New("event title is required")
	}

	event.UserID = userID

	return s.repo.Create(event)
}

func (s *eventService) GetAllEvents(userID uint) ([]models.Event, error) {
	return s.repo.GetAll(userID)
}

func (s *eventService) DeleteEvent(id uint) error {
	return s.repo.Delete(id)
}

func (s *eventService) GetEventStats(eventID uint) (map[string]int64, error) {
	total, _ := s.guestRepo.CountByStatus(eventID, "")
	accepted, _ := s.guestRepo.CountByStatus(eventID, "Accepted")
	declined, _ := s.guestRepo.CountByStatus(eventID, "Declined")
	pending, _ := s.guestRepo.CountByStatus(eventID, "Pending")

	return map[string]int64{
		"total":    total,
		"accepted": accepted,
		"declined": declined,
		"pending":  pending,
	}, nil
}
