package in_memory

import (
	"company-microservice/internals"
	company "company-microservice/pkg/Company"
	"errors"
	"time"
)

type InMemoryCompanyRepository struct {
	Companies []company.Company
}

func NewInMemoryCompanyRepository() *InMemoryCompanyRepository {
	companyMockData := []company.Company{
		{
			BaseCompany: company.BaseCompany{
				Name:              "Amazon",
				Description:       "Empresa de paqueteria a casa",
				AmountofEmployees: 500,
				Registered:        true,
				Type:              company.Corporations,
			},
			ID:      "013a7a75-b3d8-4e47-8842-8bb2541438f2",
			Created: internals.CreateDateAsDateTime(),
			Updated: internals.CreateDateAsDateTime(),
		},
	}

	return &InMemoryCompanyRepository{
		Companies: companyMockData,
	}
}

// Implementation

func (repo *InMemoryCompanyRepository) Search() ([]company.Company, error) {
	return repo.Companies, nil
}

func (repo *InMemoryCompanyRepository) Find(id string) (company.Company, error) {
	for _, company := range repo.Companies {
		if company.ID == id {
			return company, nil
		}
	}

	return company.Company{}, errors.New("company not found")
}

func (repo *InMemoryCompanyRepository) Create(data company.Company) (company.Company, error) {
	for _, company := range repo.Companies {
		if company.ID == data.ID || company.Name == data.Name {
			return company, errors.New("company already exists")
		}
	}

	newCompanys := append(repo.Companies, data)

	repo.Companies = newCompanys

	return data, nil
}

func (repo *InMemoryCompanyRepository) Update(id string, data company.Company) (company.Company, error) {
	newCompanies := make([]company.Company, len(repo.Companies))
	copy(newCompanies, repo.Companies)

	var index = -1

	for i, company := range newCompanies {
		if company.ID == id {
			index = i
			break
		}
	}

	if index < 0 {
		return company.Company{}, errors.New("do not exist a company with that id")
	}

	newCompanies[index].Name = data.Name
	newCompanies[index].Description = data.Description
	newCompanies[index].AmountofEmployees = data.AmountofEmployees
	newCompanies[index].Registered = data.Registered
	newCompanies[index].Type = data.Type
	newCompanies[index].Updated = time.Now().Format(time.DateTime)

	exist := false

	for i, company := range newCompanies {
		if i == index {
			continue
		}

		if company.Name == data.Name {
			exist = true
			break
		}
	}

	if exist {
		return company.Company{}, errors.New("company already exists")
	}

	repo.Companies = newCompanies

	return newCompanies[index], nil
}

func (repo *InMemoryCompanyRepository) Delete(id string) error {
	newCompanies := make([]company.Company, len(repo.Companies))
	copy(newCompanies, repo.Companies)

	var index = -1

	for i, company := range newCompanies {
		if company.ID == id {
			index = i
			break
		}
	}

	if index < 0 {
		return errors.New("do not exist a company with that id")
	}

	newCompanies = append(newCompanies[:index], newCompanies[index+1:]...)

	repo.Companies = newCompanies

	return nil
}
