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