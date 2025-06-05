package auth

import (
	"github.com/bryanium-templates/go-clean-architecture/models"
	"gorm.io/gorm"
)

// Repository Layer Initializer
type Repository struct {
	db *gorm.DB
}

func NewRepository (db *gorm.DB) *Repository {
	return &Repository{db : db}
}

var user models.User

// Repository Layer Methods
func (r *Repository) FindByEmail(email string) (*models.User, error) {
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil // User not found
		}
		return nil, err // Other database error		
	}
	return &user, nil 
}

func (r *Repository) CreateUser(user *models.User) (*models.User, error) {
    if err := r.db.Create(user).Error; err != nil {
        return nil, err
    }
    return user, nil
}