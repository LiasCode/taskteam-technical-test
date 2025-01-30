package userservice

type MockUserService struct {
}

func NewMockUserService() *MockUserService {
	return &MockUserService{}
}

func (s *MockUserService) ValidateUser(username string, password string) bool {
	return false
}

func (s *MockUserService) IsUserTokenValid(token string) bool {
	return true
}
