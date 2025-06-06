package auth

import (
	"github.com/google/uuid"

	"github.com/bryanium-templates/go-clean-architecture/models"
	"github.com/bryanium-templates/go-clean-architecture/utils"
)

// Service Layer Initializer
type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

// Service Layer Methods
func (s *Service) SignUp (req SignUpRequest) (*models.User, string, error) {
	if req.Email == "" || req.Password == "" || req.Username == "" {
		return nil, "", utils.ErrorHandler(400,"username, password, and email is required", nil)
	}

	if req.Password != req.ConfirmPassword {
		return nil, "", utils.ErrorHandler(400, "passwords doesn't match", nil)
	}

	existingUser, err := s.repo.FindByEmail(req.Email)
    if err != nil {
        return nil, "", utils.ErrorHandler(500, "failed to verify if user exists", err)
    }
    if existingUser != nil {
        return nil, "", utils.ErrorHandler(400,"user with email %s already exists", err)
    }

	hashedPassword,err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, "", utils.ErrorHandler(500, "failed to hash the password", err)
	}

	userId := uuid.New()

	newUser := &models.User{
		ID: userId,
		Username: req.Username,
		Email: req.Email,
		Password: hashedPassword,
		ProfilePicture: "",
	}

	createdUser, err := s.repo.CreateUser(newUser)
    if err != nil {
        return nil, "", utils.ErrorHandler(500,"failed to save user: %v", err)
    }

	token, err := utils.GenerateJWT(createdUser)
    if err != nil {
        return nil, "", utils.ErrorHandler(500,"failed to generate token: %v", err)
    }

	return  createdUser, token, nil
}
