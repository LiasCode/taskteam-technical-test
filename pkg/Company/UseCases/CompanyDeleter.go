package usecases

import (
	company "company-microservice/pkg/Company"
	eventbroker "company-microservice/pkg/EventBroker"
)

func CompanyDeleter(repo company.CompanyRepository, event_broker eventbroker.EventBroker, id string) error {

	targetCompany, err := repo.Find(id)

	if err != nil {
		return err
	}

	err = repo.Delete(targetCompany.ID)

	event_broker.Publish("company.deleted", targetCompany.ID)

	return err
}
