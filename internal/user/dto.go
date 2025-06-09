package user

type UpdateUserRequest struct {
	Username string 
	Email string 
}

type UpdateProfilePicture struct {
	ProfilePicture string `form:"profile_picture"`
}