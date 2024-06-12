package users

import "github.com/tryoasnafi/users/common"

var (
	service *UserService
)

type Service interface {
	GetAllUsers() ([]UserResponse, error)
	GetUserById(id uint) (UserResponse, error)
	CreateUser(userReq CreateUserRequest) (User, error)
	UpdateUser(id uint, user UpdateUserRequest) (User, error)
	DeleteUser(id uint) error
}

type UserService struct {
	repo Repository
}

func NewUserService(repo Repository) UserService {
	return UserService{repo: repo}
}

func (srv UserService) GetAllUsers() ([]UserResponse, error) {
	users, err := srv.repo.All()
	if err != nil {
		return []UserResponse{}, err
	}
	responses := make([]UserResponse, 0, len(users))
	for _, user := range users {
		responses = append(responses, UserToResponse(user))
	}
	return responses, nil
}

func (srv UserService) GetUserById(id uint) (UserResponse, error) {
	return UserResponse{}, common.ErrNotImplemented
}

func (srv UserService) CreateUser(userReq CreateUserRequest) (User, error) {
	user := UserFromCreateRequest(userReq)
	if err := srv.repo.Add(&user); err != nil {
		return User{}, err
	}
	return user, nil
}

func (srv UserService) UpdateUser(id uint, userReq UpdateUserRequest) (User, error) {
	user := UserFromUpdateRequest(userReq)
	user.ID = id
	if err := srv.repo.Update(&user); err != nil {
		return User{}, err
	}
	return user, nil
}
func (srv UserService) DeleteUser(id uint) error {
	return common.ErrNotImplemented
}
