package usecases

import (
	company "company-microservice/pkg/Company"
	eventbroker "company-microservice/pkg/EventBroker"
)

func CompanyUpdater(repo company.CompanyRepository, event_broker eventbroker.EventBroker, data company.BaseCompany, id string) error {
	targetCompany, err := repo.Find(id)

	if err != nil {
		return err
	}

	targetCompany.Name = data.Name
	targetCompany.Description = data.Description
	targetCompany.AmountofEmployees = data.AmountofEmployees
	targetCompany.Type = data.Type
	targetCompany.Registered = data.Registered

	if err := targetCompany.IsValid(); err != nil {
		return err
	}

	_, repoErrorResponse := repo.Update(targetCompany.ID, targetCompany)

	event_broker.Publish("company.updated", targetCompany.ID)

	return repoErrorResponse
}
