package usecases

import company "company-microservice/pkg/Company"

func CompanyFinder(repo company.CompanyRepository, id string) (company.Company, error) {
	repoResponse, err := repo.Find(id)

	if err != nil {
		return company.Company{}, err
	}

	return repoResponse, nil
}
