package repository

import (
	"simple_bank/models"
)

type UserRepository interface {
	InsertUser(u *models.User) (*models.User, error)
	GetAllUser() ([]*models.User, error)
}

type userRepository struct {
	DB []*models.User
}

func NewUserRepository() UserRepository {
	var DB []*models.User
	return &userRepository{
		DB: DB,
	}
}

func (s *userRepository) InsertUser(u *models.User) (*models.User, error) {
	user := &models.User{
		Username: u.Username,
		FullName: u.FullName,
		Email:    u.Email,
	}

	s.DB = append(s.DB, user)

	return user, nil
}

func (s *userRepository) GetAllUser() ([]*models.User, error) {

	return s.DB, nil
}
