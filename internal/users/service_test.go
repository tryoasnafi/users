package users

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var mockRepoSequenceID = 4
var mockRepoDB = []User{
	{
		ID: 1,
		FirstName: "John",
		LastName: "Doe",
		Address: "Galaxy Street",
		DOB: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
		Email: "john@example.com",
		PhoneNumber: "628123456789",
	},
	{
		ID: 2,
		FirstName: "Doe",
		LastName: "John",
		Address: "Galaxy Street 2",
		DOB: time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
		Email: "john2@example.com",
		PhoneNumber: "628123456788",
	},
	{
		ID: 3,
		FirstName: "Lily",
		LastName: "Whood",
		Address: "World Tree Street",
		DOB: time.Date(1998, time.January, 1, 0, 0, 0, 0, time.UTC),
		Email: "lilywhitey@example.com",
		PhoneNumber: "628123456780",
	},
}

type MockRepository struct {}

func (repo MockRepository) All() ([]User, error) {
	return mockRepoDB, nil
}
func (repo MockRepository) FindById(id uint) (User, error) {
	var user User
	for _, u := range mockRepoDB {
		if u.ID == id {
			user = u
			break
		}
	}
	if user.ID == 0 {
		return user, ErrUserNotFound
	}
	return user, nil
}
func (repo MockRepository) Add(user *User) error {
	isEmailExists := false
	for _, u := range mockRepoDB {
		if u.Email == user.Email {
			isEmailExists = true
			break
		}
	}
	if isEmailExists {
		return errors.New("email already exists")
	}
	mockRepoDB = append(mockRepoDB, *user)
	return nil
}
func (repo MockRepository) Update(user *User) error {
	return nil
}
func (repo MockRepository) Delete(id uint) error {
	return nil
}

func TestGetAllUsersService(t *testing.T) {
	tests := []struct{
		name string
		want []User
		wantErr error
	}{
		{
			name: "get all users record",
			want: mockRepoDB,
			wantErr: nil,
		},
	}
	srv := UserService{repo: MockRepository{}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := srv.GetAllUsers()
			assert.NoError(t, err)
			assert.Equal(t, len(tt.want), len(got))
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestGetUserByIdService(t *testing.T) {
	tests := []struct{
		name string
		want User
		wantErr error
	}{
		{
			name: "get user id 1",
			want: mockRepoDB[0],
			wantErr: nil,
		},
		{
			name: "get user id 200",
			want: User{},
			wantErr: ErrUserNotFound,
		},
	}
	srv := UserService{repo: MockRepository{}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := srv.GetUserById(tt.want.ID)
			assert.ErrorIs(t, err, tt.wantErr)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestCreateUserService(t *testing.T) {
	tests := []struct{
		name string
		user CreateUserRequest
		wantErr error
	}{
		{
			name: "create new user",
			user: CreateUserRequest{
				Email: "mimi@example.com",
				FirstName: "Miri",
				LastName: "Miri",
				DOB: time.Now(),
				PhoneNumber: "123",
				Address: "OK street",
			},
			wantErr: nil,
		},
	}
	srv := UserService{repo: MockRepository{}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := srv.CreateUser(tt.user)
			if assert.NoError(t, err) {
				assert.Equal(t, tt.user.Email, got.Email)
				assert.Equal(t, 4, len(mockRepoDB) )
			}
		})
	}
}