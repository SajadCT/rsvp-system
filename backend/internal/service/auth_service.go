package service

import (
	"errors"
	"time"

	"rsvp-system/internal/models"
	"rsvp-system/internal/repository"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthService interface {
	RegisterUser(name, email, password string) error
	LoginUser(email, password string) (string, *models.User, error)
}

type authService struct {
	repo repository.UserRepository
}

func NewAuthService(repo repository.UserRepository) AuthService {
	return &authService{repo}
}

var jwtSecret = []byte("YOUR_SUPER_SECRET_KEY")

func (s *authService) RegisterUser(name, email, password string) error {

	if _, err := s.repo.GetByEmail(email); err == nil {
		return errors.New("email already registered")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &models.User{
		Name:     name,
		Email:    email,
		Password: string(hashed),
	}
	return s.repo.Create(user)
}

func (s *authService) LoginUser(email, password string) (string, *models.User, error) {

	user, err := s.repo.GetByEmail(email)
	if err != nil {
		return "", nil, errors.New("invalid credentials")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", nil, errors.New("invalid credentials")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", nil, err
	}

	return tokenString, user, nil
}
