package services

import (
	"errors"
	"strings"

	"forum-backend/internal/models"
	"forum-backend/internal/repositories"
)

type UserService struct{
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(username string) (*models.User, error){ 
	username = strings.TrimSpace(username)

	// Validation
	if username == ""{
		return nil, errors.New("username cannot be empty")
	}

	// Check for existing username
	existingUser, err := s.repo.GetByUsername(username)
	if err != nil {
		return nil, err
	}
	
	if existingUser != nil { 
		return nil, errors.New("username already exists")
	}

	// Construct domain object
	user := &models.User {
		Username: username,
	}

	if err := s.repo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) GetUserByID(id int) (*models.User, error){
	// Validation
	if id <= 0 {
		return nil, errors.New("id cannot be less than 1")
	}

	// Check if user exists
	user, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}
	
	if user == nil { 
		return nil, errors.New("user does not exist")
	}

	return user, nil
}

func (s *UserService) GetUserByUsername(username string) (*models.User, error){
	// Validation
	if username == ""{
		return nil, errors.New("username cannot be empty")
	}

	// Check if user exists
	user, err := s.repo.GetByUsername(username)
	if err != nil {
		return nil, err
	}
	
	if user == nil { 
		return nil, errors.New("user does not exist")
	}

	return user, nil
}
