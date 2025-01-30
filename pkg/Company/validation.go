package company

import (
	"company-microservice/internals"
	"errors"
	"time"

	"github.com/google/uuid"
)

func (c *Company) IsValid() error {
	if c == nil {
		return errors.New("Company can't be null")
	}

	if len(c.Name) < 1 || len(c.Name) > 15 {
		return errors.New("name length should be between 1 and 15")
	}

	if len(c.Description) > 3000 {
		return errors.New("description max length is 3000")
	}

	if c.AmountofEmployees < 0 {
		return errors.New("amount of employees min is 0")
	}

	if !IsCompanyTypeValid(c.Type) {
		return errors.New("company type is not valid")
	}

	if uuid.Validate(c.ID) != nil {
		return errors.New("id is not valid UUID")
	}

	if !internals.IsDateValue(c.Created, time.DateTime) {
		return errors.New("created date is not valid")
	}

	if !internals.IsDateValue(c.Updated, time.DateTime) {
		return errors.New("updated date is not valid")
	}

	return nil
}

func IsCompanyTypeValid(t CompanyType) bool {
	switch t {
	case Cooperative, NonProfit, Corporations, SoleProprietorship:
		return true
	default:
		return false
	}
}
