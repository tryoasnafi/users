package users

func UserToResponse(user User) UserResponse {
	return UserResponse{
		ID: int(user.ID),
		FirstName: user.FirstName,
		LastName: user.LastName,
		Address: user.Address,
		DOB: user.DOB,
		Email: user.Email,
		PhoneNumber: user.PhoneNumber,
	}
}

func UserFromCreateRequest(user CreateUserRequest) User {
	return User{
		FirstName: user.FirstName,
		LastName: user.LastName,
		Address: user.Address,
		Email: user.Email,
		DOB: user.DOB,
		PhoneNumber: user.PhoneNumber,
	}
}

func UserFromUpdateRequest(user UpdateUserRequest) User {
	return User{
		FirstName: user.FirstName,
		LastName: user.LastName,
		Address: user.Address,
		Email: user.Email,
		DOB: user.DOB,
		PhoneNumber: user.PhoneNumber,
	}
}