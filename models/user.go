package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`

	Username string
	Email string `gorm:"unique;not null"`
	Password string `gorm:"not null"`

	ProfilePicture string `gorm:"default:https://fallback.com/profile.png"`


	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time	`gorm:"autoUpdateTime"`
}