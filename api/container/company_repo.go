package container

import (
	in_memory "company-microservice/api/repositories/InMemory"
	postgresql "company-microservice/api/repositories/Postgresql"
	company "company-microservice/pkg/Company"
)

func (c *Container) GetCompanyRepository() company.CompanyRepository {
	if c.companyRepositoryInstance != nil {
		return c.companyRepositoryInstance
	}

	if c.IsTesting {
		inMemoryCompanyRepository := in_memory.NewInMemoryCompanyRepository()

		if inMemoryCompanyRepository == nil {
			panic("InMemoryCompanyRepository is nil")
		}

		c.companyRepositoryInstance = inMemoryCompanyRepository
		return c.companyRepositoryInstance
	}

	postgresqlCompanyRepository := postgresql.NewPostgresCompanyRepository(c.GetPostgresqlDB())

	if postgresqlCompanyRepository == nil {
		panic("postgresqlCompanyRepository is nil")
	}

	c.companyRepositoryInstance = postgresqlCompanyRepository

	return c.companyRepositoryInstance
}
