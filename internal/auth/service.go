package auth

import (
	"fmt"
	"time"

	"github.com/bryanium-templates/go-clean-architecture/models"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
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
		return nil, "", fmt.Errorf("email, password and username are required")
	}

	existingUser, err := s.repo.FindByEmail(req.Email)
    if err != nil {
        return nil, "", fmt.Errorf("failed to check user existence: %v", err)
    }
    if existingUser != nil {
        return nil, "", fmt.Errorf("user with email %s already exists", req.Email)
    }

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
    if err != nil {
        return nil, "", fmt.Errorf("failed to hash password: %v", err)
    }

	newUser := &models.User{
		Username: req.Username,
		Email: req.Email,
		Password: string(hashedPassword),
		ProfilePicture: "",
	}

	createdUser, err := s.repo.CreateUser(newUser)
    if err != nil {
        return nil, "", fmt.Errorf("failed to save user: %v", err)
    }

	token, err := s.generateJWT(createdUser)
    if err != nil {
        return nil, "", fmt.Errorf("failed to generate token: %v", err)
    }

	return  createdUser, token, nil
}


func (s *Service) generateJWT(user *models.User) (string, error) {
    // Create JWT claims
    claims := jwt.MapClaims{
        "sub": user.ID,         // subject, usually the user ID
        "email": user.Email,    // additional user info
        "username": user.Username,
        "exp": time.Now().Add(time.Hour * 24).Unix(), // token expiration time
    }

    // Generate the token with your secret key
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString([]byte("your-secret-key"))
    if err != nil {
        return "", err
    }

    return tokenString, nil
}