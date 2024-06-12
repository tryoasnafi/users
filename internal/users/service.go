package users

var (
	service *UserService
)

type Service interface {
	GetAllUsers() ([]User, error)
	GetUserById(id uint) (User, error)
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

func (srv UserService) GetAllUsers() ([]User, error) {
	users, err := srv.repo.All()
	if err != nil {
		return []User{}, err
	}
	return users, nil
}

func (srv UserService) GetUserById(id uint) (User, error) {
	user, err := srv.repo.FindById(id)
	if err != nil {
		return User{}, err
	}
	return user, nil
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
	return srv.repo.Delete(id)
}
