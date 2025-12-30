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

func (m *MockEventRepo) GetAll(userID uint) ([]models.Event, error) {
	args := m.Called(userID)
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

type MockUserRepo struct {
	mock.Mock
}

func (m *MockUserRepo) Create(user *models.User) error {
	return m.Called(user).Error(0)
}

func (m *MockUserRepo) GetByEmail(email string) (*models.User, error) {
	args := m.Called(email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockUserRepo) GetByID(id uint) (*models.User, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.User), args.Error(1)
}
