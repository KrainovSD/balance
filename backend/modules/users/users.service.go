package users

type UsersService struct {
	UsersProvider UsersProvider
}

func (r *UsersService) GetUser(userID int) (User, error) {
	return r.UsersProvider.GetUserById(userID)
}
