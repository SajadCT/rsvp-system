package service

import (
	"errors"
	"rsvp-system/internal/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInviteGuest_Success(t *testing.T) {
	mockEventRepo := new(MockEventRepo)
	mockGuestRepo := new(MockGuestRepo)
	guestService := NewGuestService(mockGuestRepo, mockEventRepo)

	newGuest := &models.Guest{Name: "Alice", EventID: 1}
	mockEvent := &models.Event{Title: "Existing Event"}

	// Expectation 1: Check Event Exists
	mockEventRepo.On("GetByID", uint(1)).Return(mockEvent, nil)

	// Expectation 2: Create Guest
	mockGuestRepo.On("Create", newGuest).Return(nil)

	err := guestService.InviteGuest(newGuest)

	assert.NoError(t, err)
	mockEventRepo.AssertExpectations(t)
	mockGuestRepo.AssertExpectations(t)
}

func TestInviteGuest_EventNotFound(t *testing.T) {
	mockEventRepo := new(MockEventRepo)
	mockGuestRepo := new(MockGuestRepo)
	guestService := NewGuestService(mockGuestRepo, mockEventRepo)

	newGuest := &models.Guest{Name: "Bob", EventID: 999}

	// Expectation: Event not found
	mockEventRepo.On("GetByID", uint(999)).Return(nil, errors.New("db error"))

	err := guestService.InviteGuest(newGuest)

	assert.Error(t, err)
	assert.Equal(t, "event not found", err.Error())
}

func TestRSVPGuest_Success(t *testing.T) {
	mockEventRepo := new(MockEventRepo)
	mockGuestRepo := new(MockGuestRepo)
	guestService := NewGuestService(mockGuestRepo, mockEventRepo)

	guestID := uint(50)
	status := "Declined"

	// Mock existing guest
	mockGuest := &models.Guest{Name: "Sam", Status: "Pending"}

	// Expectation 1: Find Guest
	mockGuestRepo.On("GetByID", guestID).Return(mockGuest, nil)

	// Expectation 2: Update
	mockGuestRepo.On("UpdateStatus", mockGuest, status).Return(nil)

	err := guestService.RSVPGuest(guestID, status)

	assert.NoError(t, err)
	mockGuestRepo.AssertExpectations(t)
}
