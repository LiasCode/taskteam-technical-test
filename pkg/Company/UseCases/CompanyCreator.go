package usecases

import (
	company "company-microservice/pkg/Company"
	eventbroker "company-microservice/pkg/EventBroker"
)

func CompanyCreator(repo company.CompanyRepository, event_broker eventbroker.EventBroker, data company.BaseCompany) (company.Company, error) {
	newCompany, err := company.NewCompany(
		data.Name,
		data.Description,
		data.AmountofEmployees,
		data.Registered,
		data.Type,
	)

	if err != nil {
		return company.Company{}, err
	}

	if err := newCompany.IsValid(); err != nil {
		return company.Company{}, err
	}

	repoResponse, err := repo.Create(*newCompany)

	if err != nil {
		return company.Company{}, err
	}

	event_broker.Publish("company.created", newCompany.ID)

	return repoResponse, nil
}
