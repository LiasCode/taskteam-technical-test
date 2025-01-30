package container

import (
	company "company-microservice/pkg/Company"
	eventbroker "company-microservice/pkg/EventBroker"
	user "company-microservice/pkg/User"
	"database/sql"
)

type Container struct {
	IsTesting                 bool
	companyRepositoryInstance company.CompanyRepository
	userService               user.UserService
	eventBroker               eventbroker.EventBroker
	postgresqlDB              *sql.DB
}

var ContainerInstance = &Container{
	IsTesting: false,
}
