package users

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/tryoasnafi/users/internal/validation"
)

var mockSequenceID = 3
var mockDB = []User{
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
}

type MockUserService struct{}

func (srv MockUserService) GetAllUsers() ([]User, error) {
	return mockDB, nil
}
func (srv MockUserService) GetUserById(id uint) (User, error) {
	var user User
	for _, u := range mockDB {
		if u.ID == id {
			user = u
			break
		}
	}
	if user.ID == 0 {
		log.Println("Got hit with id", id)
		return user, ErrUserNotFound
	}
	return user, nil
}
func (srv MockUserService) CreateUser(userReq CreateUserRequest) (User, error) {
	user := UserFromCreateRequest(userReq)
	user.ID = uint(mockSequenceID)
	mockSequenceID++
	mockDB = append(mockDB, user)
	return user, nil
}
func (srv MockUserService) UpdateUser(id uint, userReq UpdateUserRequest) (User, error) {
	return User{}, nil
}
func (srv MockUserService) DeleteUser(id uint) error {
	return nil
}

func TestGetAllUsers(t *testing.T) {
	tests := []struct{
		id int
		name string
		expectedHTTPStatus int
		expectedResponse []UserResponse
	}{
		{
			name: "get all users endpoint",
			expectedHTTPStatus: http.StatusOK,
			expectedResponse: ListUserToResponse(mockDB),
		},
	}

	e := echo.New()
	h := UserHandler{service: MockUserService{}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/api/v1/", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/users")
			h.GetAllUsers(c)
		
			got := []UserResponse{}
			json.Unmarshal(rec.Body.Bytes(), &got)
			assert.Equal(t, tt.expectedResponse, got)
			assert.Equal(t, tt.expectedHTTPStatus, rec.Code)
		})
	}
}

func TestGetUserByID(t *testing.T) {
	tests := []struct{
		id int
		expectedHTTPStatus int
		expectedResponse UserResponse
	}{
		{
			id: 1,
			expectedHTTPStatus: http.StatusOK,
			expectedResponse: UserToResponse(mockDB[0]),
		},
		{
			id: 2,
			expectedHTTPStatus: http.StatusOK,
			expectedResponse: UserToResponse(mockDB[1]),
		},
		{
			id: 999,
			expectedHTTPStatus: http.StatusNotFound,
			expectedResponse: UserResponse{},
		},
	}

	e := echo.New()
	h := UserHandler{service: MockUserService{}}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("ID %d", tt.id), func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/api/v1/", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/users/:id")
			c.SetParamNames("id")
			c.SetParamValues(fmt.Sprintf("%d", tt.id))
			h.GetUserById(c)
		
			got := UserResponse{}
			json.Unmarshal(rec.Body.Bytes(), &got)
			assert.Equal(t, tt.expectedResponse, got)
			assert.Equal(t, tt.expectedHTTPStatus, rec.Code)
		})
	}
}

func TestCreateUser(t *testing.T) {
	newUserMick := CreateUserRequest{
		Username:    "mickmous",
		FirstName:   "Mick",
		LastName:    "Mous",
		Address:     "Galaxy Street 243",
		DOB:         time.Date(2000, time.January, 1, 0, 0, 0, 0, time.UTC),
		Email:       "mickmous@example.com",
		PhoneNumber: "6281234567889",
	}

	tests := []struct{
		name               string
		user               CreateUserRequest
		expectedHTTPStatus int
		expectedResponse   UserResponse
	}{
		{
			name:               "create a user",
			user:               newUserMick,
			expectedHTTPStatus: http.StatusCreated,
			expectedResponse: func() UserResponse {
				user := UserToResponse(UserFromCreateRequest(newUserMick))
				user.ID = mockSequenceID
				return user
			}(),
		},
	}

	e := echo.New()
	e.Validator = &validation.CustomValidator{Validator: validator.New()}
	h := UserHandler{service: MockUserService{}}

	for _, tt := range tests {
		currentRecords := len(mockDB)
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			json.NewEncoder(&buf).Encode(tt.user)
			req := httptest.NewRequest(http.MethodPost, "/api/v1/", &buf)
			rec := httptest.NewRecorder()
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			c := e.NewContext(req, rec)
			c.SetPath("/users")
			err := h.CreateUser(c)
			if err != nil {
				t.Log(err)
			}

			got := UserResponse{}
			json.Unmarshal(rec.Body.Bytes(), &got)
			assert.Equal(t, tt.expectedHTTPStatus, rec.Code)
			assert.Equal(t, tt.expectedResponse, got)
			assert.Equal(t, currentRecords+1, len(mockDB))
		})
	}
}
