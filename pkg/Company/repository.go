package company

type CompanyRepository interface {
	Search() ([]Company, error)

	Find(id string) (Company, error)

	Create(data Company) (Company, error)

	Update(id string, data Company) (Company, error)

	Delete(id string) error
}
