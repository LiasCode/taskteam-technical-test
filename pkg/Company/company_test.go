package company

import "testing"

func TestCompanyValidation(t *testing.T) {
	company, creationError := NewCompany("NombreTest", "DescriptionTest", 10, false, Corporations)

	t.Run("Company Creation", func(t *testing.T) {

		if creationError != nil {
			t.Errorf("Company creation should not be and error")
		}

		if company.Name != "NombreTest" {
			t.Errorf("Company Name should be `NombreTest`")
		}

		if company.Description != "DescriptionTest" {
			t.Errorf("Company Description should be `DescriptionTest`")
		}

		if company.AmountofEmployees != 10 {
			t.Errorf("Company AmountofEmployees should be `10`")
		}

		if company.Registered {
			t.Errorf("Company Registered should be `false`")
		}

		if company.Type != Corporations {
			t.Errorf("Company Corporations should be `Corporations`")
		}
	})

	t.Run("Company Validation Method", func(t *testing.T) {
		scope_company := company

		if scope_company.IsValid() != nil {
			t.Errorf("Company should be valid")
		}

		scope_company.ID = "1"

		if scope_company.IsValid() == nil {
			t.Errorf("Company should not be valid")
		}

		scope_company.ID = company.ID
		scope_company.Name = "123456789101112131415"

		if scope_company.IsValid() == nil {
			t.Errorf("Company should not be valid, the name max length should be 15")
		}
	})
}
