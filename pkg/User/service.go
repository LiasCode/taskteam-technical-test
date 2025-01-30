package user

type UserService interface {
	ValidateUser(username string, password string) bool
	IsUserTokenValid(token string) bool
}
