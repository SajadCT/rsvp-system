package service

import (
	"rsvp-system/internal/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateEvent_Success(t *testing.T) {
	mockEventRepo := new(MockEventRepo)
	mockGuestRepo := new(MockGuestRepo)
	eventService := NewEventService(mockEventRepo, mockGuestRepo)

	userID := uint(1)
	newEvent := &models.Event{Title: "Service Test Party"}

	mockEventRepo.On("Create", newEvent).Return(nil)

	err := eventService.CreateEvent(newEvent, userID)

	assert.NoError(t, err)
	assert.Equal(t, userID, newEvent.UserID)
	mockEventRepo.AssertExpectations(t)
}

func TestGetEventStats(t *testing.T) {
	mockEventRepo := new(MockEventRepo)
	mockGuestRepo := new(MockGuestRepo)
	eventService := NewEventService(mockEventRepo, mockGuestRepo)

	eventID := uint(100)

	mockGuestRepo.On("CountByStatus", eventID, "").Return(20, nil)
	mockGuestRepo.On("CountByStatus", eventID, "Accepted").Return(15, nil)
	mockGuestRepo.On("CountByStatus", eventID, "Declined").Return(2, nil)
	mockGuestRepo.On("CountByStatus", eventID, "Pending").Return(3, nil)

	stats, err := eventService.GetEventStats(eventID)

	assert.NoError(t, err)
	assert.Equal(t, int64(20), stats["total"])
	assert.Equal(t, int64(15), stats["accepted"])
}
