package auth

import "github.com/google/uuid"

type SignUpRequest struct {
	Username string
	Email string 
	Password string
	ConfirmPassword string
}

type SignUpResponse struct {
	ID uuid.UUID

}