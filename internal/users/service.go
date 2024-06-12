package users

import "github.com/tryoasnafi/users/common"

var (
	service *UserService
)

type Service interface {
	GetAllUsers() ([]UserResponse, error)
	GetUserById(id uint) (UserResponse, error)
	CreateUser(user CreateUserRequest) (UserResponse, error)
	UpdateUser(user UpdateUserRequest) (UserResponse, error)
	DeleteUser(id uint) error
}

type UserService struct {
	repo Repository
}

func NewUserService(repo Repository) UserService {
	return UserService{repo: repo}
}

func (srv UserService) GetAllUsers() ([]UserResponse, error) {
	return []UserResponse{}, common.ErrNotImplemented
}

func (srv UserService) GetUserById(id uint) (UserResponse, error) {
	return UserResponse{}, common.ErrNotImplemented
}

func (srv UserService) CreateUser(user CreateUserRequest) (UserResponse, error) {
	return UserResponse{}, common.ErrNotImplemented
}

func (srv UserService) UpdateUser(user UpdateUserRequest) (UserResponse, error) {
	return UserResponse{}, common.ErrNotImplemented
}
func (srv UserService) DeleteUser(id uint) error {
	return common.ErrNotImplemented
}
