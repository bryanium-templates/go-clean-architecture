package user

// Service Layer Initializer
type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

// Service Layer Methods
func (s *Service) UpdateUser (req UpdateUserRequest) {

}

func (s *Service) UpdateProfilePicture (req UpdateProfilePicture) {

}

func (s *Service) DeleteUser (userId string) {

}

func (s *Service) GetAllUsers () {

}