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
	user := UserFromUpdateRequest(userReq)
	user.ID = id
	isFound := false
	for _, u := range mockDB {
		if u.ID == id {
			u = user
			isFound = true
			break
		}
	}
	if !isFound {
		log.Println("Right here")
		return User{}, ErrUserNotFound
	}
	return user, nil
}
func (srv MockUserService) DeleteUser(id uint) error {
	isFound := false
	for _, u := range mockDB {
		if u.ID == id {
			isFound = true
			break
		}
	}
	if !isFound {
		return ErrUserNotFound
	}
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

func TestUpdateUser(t *testing.T) {
	type TestCase struct {
		name               string
		id                 uint
		user               UpdateUserRequest
		expectedHTTPStatus int
		expectedResponse   any
	}
	tests := []TestCase {
		{
			name: "update user id 1",
			id:   mockDB[0].ID,
			user: func() UpdateUserRequest {
				oldData := mockDB[0]
				return UpdateUserRequest{
					FirstName:   "New John",
					LastName:    "New Doe",
					DOB:         time.Date(2001, time.January, 1, 0, 0, 0, 0, time.UTC),
					Email:       oldData.Email,
					PhoneNumber: oldData.PhoneNumber,
					Address:     "New Address",
				}
			}(),
			expectedHTTPStatus: http.StatusOK,
			expectedResponse: func() UserResponse {
				oldData := mockDB[0]
				return UserToResponse(User{
					ID:          oldData.ID,
					FirstName:   "New John",
					LastName:    "New Doe",
					DOB:         time.Date(2001, time.January, 1, 0, 0, 0, 0, time.UTC),
					Email:       oldData.Email,
					PhoneNumber: oldData.PhoneNumber,
					Address:     "New Address",
				})
			}(),
		},
		{
			name: "update user id 100, should not found",
			id:   100,
			user: func() UpdateUserRequest {
				oldData := mockDB[0]
				return UpdateUserRequest{
					FirstName:   "New John",
					LastName:    "New Doe",
					DOB:         time.Date(2001, time.January, 1, 0, 0, 0, 0, time.UTC),
					Email:       oldData.Email,
					PhoneNumber: oldData.PhoneNumber,
					Address:     "New Address",
				}
			}(),
			expectedHTTPStatus: http.StatusNotFound,
			expectedResponse: MessageResponse{Message: ErrUserNotFound.Error()},
		},
	}

	e := echo.New()
	e.Validator = &validation.CustomValidator{Validator: validator.New()}
	h := UserHandler{service: MockUserService{}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var buf bytes.Buffer
			json.NewEncoder(&buf).Encode(tt.user)
			req := httptest.NewRequest(http.MethodPut, "/api/v1/", &buf)
			rec := httptest.NewRecorder()
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
			c := e.NewContext(req, rec)
			c.SetPath("/users/:id")
			c.SetParamNames("id")
			c.SetParamValues(fmt.Sprintf("%d", tt.id))
			if err := h.UpdateUser(c); err != nil {
				t.Log(err)
			}

			assert.Equal(t, tt.expectedHTTPStatus, rec.Code)
			
			got := tt.expectedResponse
			json.Unmarshal(rec.Body.Bytes(), &got)
			switch tt.expectedResponse.(type) {
			case UserResponse:
				if val, ok := tt.expectedResponse.(UserResponse); ok {
					assert.Equal(t, tt.expectedResponse, val)
				}
			case MessageResponse:
				if val, ok := tt.expectedResponse.(MessageResponse); ok {
					assert.Equal(t, tt.expectedResponse, val)
				}
			}
		})
	}
}

func TestDeleteUser(t *testing.T) {
	type TestCase struct {
		name               string
		id                 uint
		expectedHTTPStatus int
		expectedResponse   any
	}
	tests := []TestCase {
		{
			name: "delete user id 1",
			id:   mockDB[0].ID,
			expectedHTTPStatus: http.StatusOK,
		},
		{
			name: "delete user id 100, should not found",
			id:   100,
			expectedHTTPStatus: http.StatusNotFound,
		},
	}

	e := echo.New()
	h := UserHandler{service: MockUserService{}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodPut, "/api/v1/", nil)
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/users/:id")
			c.SetParamNames("id")
			c.SetParamValues(fmt.Sprintf("%d", tt.id))
			if err := h.DeleteUser(c); err != nil {
				t.Log(err)
			}

			assert.Equal(t, tt.expectedHTTPStatus, rec.Code)
		})
	}
}
