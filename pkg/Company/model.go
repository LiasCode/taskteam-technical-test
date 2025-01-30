package company

import (
	"time"

	"github.com/google/uuid"
)

type CompanyType string

const (
	Corporations       CompanyType = "Corporations"
	NonProfit          CompanyType = "NonProfit"
	Cooperative        CompanyType = "Cooperative"
	SoleProprietorship CompanyType = "SoleProprietorship"
)

type BaseCompany struct {
	Name              string      `json:"name"`
	Description       string      `json:"description"`
	AmountofEmployees int         `json:"amount_of_employees"`
	Registered        bool        `json:"registered"`
	Type              CompanyType `json:"type"`
}

type Company struct {
	BaseCompany
	ID      string `json:"id"`
	Created string `json:"created"`
	Updated string `json:"updated"`
}

func NewCompany(
	name string,
	description string,
	amountofEmployees int,
	registered bool,
	company_type CompanyType,
) (*Company, error) {

	var newCompany = &Company{
		BaseCompany: BaseCompany{
			Name:              name,
			Description:       description,
			AmountofEmployees: amountofEmployees,
			Registered:        registered,
			Type:              company_type,
		},
		ID:      uuid.New().String(),
		Created: time.Now().Format(time.DateTime),
		Updated: time.Now().Format(time.DateTime),
	}

	return newCompany, nil
}
