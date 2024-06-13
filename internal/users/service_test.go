package users

import (
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
	return User{}, nil
}
func (repo MockRepository) Add(user *User) error {
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
