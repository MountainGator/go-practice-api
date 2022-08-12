package services

type UserService interface {
	CreateUser(*models.User) error
	GetUser(*string) (*models.User, error)
	// get all is a slice that contains our user objects
	GetAll() ([]*models.User, error)
	UpdateUser(*models.User) error
	// pass in username to delete user
	DeleteUser(*string) error
}