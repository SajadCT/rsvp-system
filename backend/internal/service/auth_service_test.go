package service

import (
	"errors"
	"testing"

	"rsvp-system/internal/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func TestRegisterUser_Success(t *testing.T) {

	mockUserRepo := new(MockUserRepo)
	authService := NewAuthService(mockUserRepo)

	email := "new@test.com"
	password := "password123"

	mockUserRepo.On("GetByEmail", email).Return(nil, gorm.ErrRecordNotFound)

	mockUserRepo.On("Create", mock.MatchedBy(func(u *models.User) bool {
		return u.Email == email && u.Password != password
	})).Return(nil)

	err := authService.RegisterUser("New User", email, password)

	assert.NoError(t, err)
	mockUserRepo.AssertExpectations(t)
}

func TestRegisterUser_EmailAlreadyExists(t *testing.T) {
	mockUserRepo := new(MockUserRepo)
	authService := NewAuthService(mockUserRepo)

	email := "existing@test.com"
	existingUser := &models.User{Email: email}

	mockUserRepo.On("GetByEmail", email).Return(existingUser, nil)

	err := authService.RegisterUser("User", email, "pass")

	assert.Error(t, err)
	assert.Equal(t, "email already registered", err.Error())
	mockUserRepo.AssertNotCalled(t, "Create")
}

func TestLoginUser_Success(t *testing.T) {
	mockUserRepo := new(MockUserRepo)
	authService := NewAuthService(mockUserRepo)

	email := "valid@test.com"
	password := "secret123"

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	mockUser := &models.User{
		Model:    gorm.Model{ID: 1},
		Name:     "Valid User",
		Email:    email,
		Password: string(hashedPassword),
	}

	mockUserRepo.On("GetByEmail", email).Return(mockUser, nil)

	token, user, err := authService.LoginUser(email, password)

	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, uint(1), user.ID)
	assert.NotEmpty(t, token, "JWT Token should be generated")
}

func TestLoginUser_WrongPassword(t *testing.T) {
	mockUserRepo := new(MockUserRepo)
	authService := NewAuthService(mockUserRepo)

	email := "valid@test.com"
	realPassword := "correct_password"
	wrongPassword := "wrong_password"

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(realPassword), bcrypt.DefaultCost)

	mockUser := &models.User{
		Email:    email,
		Password: string(hashedPassword),
	}

	mockUserRepo.On("GetByEmail", email).Return(mockUser, nil)

	token, user, err := authService.LoginUser(email, wrongPassword)

	assert.Error(t, err)
	assert.Equal(t, "invalid credentials", err.Error())
	assert.Empty(t, token)
	assert.Nil(t, user)
}

func TestLoginUser_UserNotFound(t *testing.T) {
	mockUserRepo := new(MockUserRepo)
	authService := NewAuthService(mockUserRepo)

	email := "ghost@test.com"

	mockUserRepo.On("GetByEmail", email).Return(nil, errors.New("record not found"))

	token, user, err := authService.LoginUser(email, "any_pass")

	assert.Error(t, err)
	assert.Equal(t, "invalid credentials", err.Error())
	assert.Empty(t, token)
	assert.Nil(t, user)
}
