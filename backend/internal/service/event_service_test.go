package service

import (
	"rsvp-system/internal/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateEvent_Success(t *testing.T) {
	// 1. Setup Mocks
	mockEventRepo := new(MockEventRepo)
	mockGuestRepo := new(MockGuestRepo)

	// Inject mocks into the real service
	eventService := NewEventService(mockEventRepo, mockGuestRepo)

	newEvent := &models.Event{Title: "Service Test Party"}

	// 2. Expectation
	mockEventRepo.On("Create", newEvent).Return(nil)

	// 3. Execution
	err := eventService.CreateEvent(newEvent)

	// 4. Assertion
	assert.NoError(t, err)
	mockEventRepo.AssertExpectations(t)
}

func TestGetEventStats(t *testing.T) {
	mockEventRepo := new(MockEventRepo)
	mockGuestRepo := new(MockGuestRepo)
	eventService := NewEventService(mockEventRepo, mockGuestRepo)

	eventID := uint(100)

	// Expectations
	mockGuestRepo.On("CountByStatus", eventID, "").Return(20, nil)
	mockGuestRepo.On("CountByStatus", eventID, "Accepted").Return(15, nil)
	mockGuestRepo.On("CountByStatus", eventID, "Declined").Return(2, nil)
	mockGuestRepo.On("CountByStatus", eventID, "Pending").Return(3, nil)

	// Execute
	stats, err := eventService.GetEventStats(eventID)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, int64(20), stats["total"])
	assert.Equal(t, int64(15), stats["accepted"])
}

func TestCreateEvent_ValidationFailure(t *testing.T) {
	mockEventRepo := new(MockEventRepo)
	mockGuestRepo := new(MockGuestRepo)
	eventService := NewEventService(mockEventRepo, mockGuestRepo)

	// Title is empty, should fail validation
	invalidEvent := &models.Event{Title: ""}

	err := eventService.CreateEvent(invalidEvent)

	assert.Error(t, err)
	assert.Equal(t, "event title is required", err.Error())
	// Ensure repo.Create was NEVER called
	mockEventRepo.AssertNotCalled(t, "Create")
}
