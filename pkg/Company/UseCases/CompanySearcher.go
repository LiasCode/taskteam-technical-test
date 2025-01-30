package usecases

import company "company-microservice/pkg/Company"

func CompanySearcher(repo company.CompanyRepository) ([]company.Company, error) {
	return repo.Search()
}
