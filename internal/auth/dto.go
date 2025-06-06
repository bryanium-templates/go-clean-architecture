package auth

type SignUpRequest struct {
	Username string
	Email string 
	Password string
	ConfirmPassword string
}

type SignInRequest struct { 
	Email string
	Password string
}