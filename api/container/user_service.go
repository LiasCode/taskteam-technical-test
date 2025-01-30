package container

import (
	userservice "company-microservice/api/services/UserService"
	user "company-microservice/pkg/User"
)

func (c *Container) GetUserService() user.UserService {
	if c.userService != nil {
		return c.userService
	}

	mock_userService := userservice.NewMockUserService()

	if mock_userService == nil {
		panic("mock_userService is nil")
	}

	c.userService = mock_userService

	return c.userService
}
