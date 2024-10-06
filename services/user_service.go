package services

import (
	"go-crud-app/models"
	"go-crud-app/config"
	"golang.org/x/crypto/bcrypt"
	"errors"
	"go-crud-app/utils"
)

type UserService interface {
	Register(user *models.User) (*models.User, error)
	Login(email, password string) (string, error)
	GetAllUsers() ([]models.User, error)
}

type userService struct {
}

func NewUserService() UserService {
	return &userService{}
}

func (s *userService) Register(user *models.User) (*models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.Password = string(hashedPassword)

	if err := config.DB.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (s *userService) Login(email, password string) (string, error) {
    var user models.User
    if err := config.DB.Where("email = ?", email).First(&user).Error; err != nil {
        return "", errors.New("invalid email or password")
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
        return "", errors.New("invalid email or password")
    }

    token, err := utils.GenerateJWT(user.Username, user.ID)
    if err != nil {
        return "", err
    }

    return token, nil
}

func (s *userService) GetAllUsers() ([]models.User, error) {
	var users []models.User
	if err := config.DB.Select("id, username, email").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
