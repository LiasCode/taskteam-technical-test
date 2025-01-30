package postgresql

import (
	company "company-microservice/pkg/Company"
	"database/sql"
	"fmt"
)

type PostgresCompanyRepository struct {
	DB *sql.DB
}

func NewPostgresCompanyRepository(db *sql.DB) *PostgresCompanyRepository {
	// Create Company Table if Not Exist
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS companies (
			id VARCHAR PRIMARY KEY,
			name VARCHAR(15) NOT NULL UNIQUE,
			description VARCHAR(3000) NOT NULL,
			amount_of_employees INTEGER NOT NULL,
			registered BOOLEAN NOT NULL,
			type VARCHAR(255) NOT NULL,
			created VARCHAR NOT NULL,
			updated VARCHAR NOT NULL
		);
	`)

	if err != nil {
		fmt.Printf("Error creating table: %v\n", err.Error())
		return nil
	}

	return &PostgresCompanyRepository{DB: db}
}

func (repo *PostgresCompanyRepository) Search() ([]company.Company, error) {
	rows, err := repo.DB.Query("SELECT * FROM companies")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var companies []company.Company

	for rows.Next() {
		var id string
		var name string
		var description string
		var amount_of_employees int
		var registered bool
		var company_type string
		var created string
		var updated string

		if err := rows.Scan(&id, &name, &description, &amount_of_employees, &registered, &company_type, &created, &updated); err != nil {
			return nil, err
		}

		companies = append(companies, company.Company{
			BaseCompany: company.BaseCompany{
				Name:              name,
				Description:       description,
				AmountofEmployees: amount_of_employees,
				Registered:        registered,
				Type:              company.CompanyType(company_type),
			},
			ID:      id,
			Created: created,
			Updated: updated,
		})
	}

	return companies, nil
}

func (repo *PostgresCompanyRepository) Find(id string) (company.Company, error) {
	query := "SELECT * FROM companies WHERE id = $1"

	row := repo.DB.QueryRow(query, id)

	var company company.Company

	err := row.Scan(
		&company.ID,
		&company.Name,
		&company.Description,
		&company.AmountofEmployees,
		&company.Registered,
		(*string)(&company.Type),
		&company.Created,
		&company.Updated,
	)

	if err != nil {
		return company, err
	}

	return company, nil
}

func (repo *PostgresCompanyRepository) Create(data company.Company) (company.Company, error) {
	query := `
		INSERT INTO companies (id, name, description, amount_of_employees, registered, type, created, updated)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id, name, description, amount_of_employees, registered, type, created, updated;
	`

	row := repo.DB.QueryRow(query, data.ID, data.Name, data.Description, data.AmountofEmployees, data.Registered, string(data.Type), data.Created, data.Updated)

	var createdCompany company.Company

	err := row.Scan(
		&createdCompany.ID,
		&createdCompany.Name,
		&createdCompany.Description,
		&createdCompany.AmountofEmployees,
		&createdCompany.Registered,
		(*string)(&createdCompany.Type),
		&createdCompany.Created,
		&createdCompany.Updated,
	)

	if err != nil {
		return company.Company{}, err
	}

	return createdCompany, nil
}

func (repo *PostgresCompanyRepository) Update(id string, data company.Company) (company.Company, error) {
	query := `
		UPDATE companies
		SET name = $2, description = $3, amount_of_employees = $4, registered = $5, type = $6, updated = $7
		WHERE id = $1
		RETURNING id, name, description, amount_of_employees, registered, type, created, updated;
	`

	row := repo.DB.QueryRow(query,
		id,
		data.Name,
		data.Description,
		data.AmountofEmployees,
		data.Registered,
		string(data.Type),
		data.Updated,
	)

	var updatedCompany company.Company

	err := row.Scan(
		&updatedCompany.ID,
		&updatedCompany.Name,
		&updatedCompany.Description,
		&updatedCompany.AmountofEmployees,
		&updatedCompany.Registered,
		(*string)(&updatedCompany.Type),
		&updatedCompany.Created,
		&updatedCompany.Updated,
	)

	if err != nil {
		return company.Company{}, err
	}

	return updatedCompany, nil
}

func (repo *PostgresCompanyRepository) Delete(id string) error {
	query := "DELETE FROM companies WHERE id = $1"

	_, err := repo.DB.Exec(query, id)

	if err != nil {
		return err
	}

	return nil
}
