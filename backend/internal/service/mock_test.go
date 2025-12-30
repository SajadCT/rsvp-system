package service

import (
	"rsvp-system/internal/models"

	"github.com/stretchr/testify/mock"
)

type MockEventRepo struct {
	mock.Mock
}

func (m *MockEventRepo) Create(event *models.Event) error {
	return m.Called(event).Error(0)
}

func (m *MockEventRepo) GetAll() ([]models.Event, error) {
	args := m.Called()
	return args.Get(0).([]models.Event), args.Error(1)
}

func (m *MockEventRepo) GetByID(id uint) (*models.Event, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Event), args.Error(1)
}

func (m *MockEventRepo) Delete(id uint) error {
	return m.Called(id).Error(0)
}

// --- MockGuestRepo ---
// Simulates the Guest Database
type MockGuestRepo struct {
	mock.Mock
}

func (m *MockGuestRepo) Create(guest *models.Guest) error {
	return m.Called(guest).Error(0)
}

func (m *MockGuestRepo) GetByEventID(id uint) ([]models.Guest, error) {
	args := m.Called(id)
	return args.Get(0).([]models.Guest), args.Error(1)
}

func (m *MockGuestRepo) GetByID(id uint) (*models.Guest, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Guest), args.Error(1)
}

func (m *MockGuestRepo) UpdateStatus(g *models.Guest, s string) error {
	return m.Called(g, s).Error(0)
}

func (m *MockGuestRepo) GetGuestWithEvent(id uint) (*models.Guest, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Guest), args.Error(1)
}

func (m *MockGuestRepo) CountByStatus(eventID uint, status string) (int64, error) {
	args := m.Called(eventID, status)
	return int64(args.Int(0)), args.Error(1)
}
