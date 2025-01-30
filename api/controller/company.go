package controller

import (
	company "company-microservice/pkg/Company"
	usecases "company-microservice/pkg/Company/UseCases"
	eventbroker "company-microservice/pkg/EventBroker"
	user "company-microservice/pkg/User"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func SearchCompanys(repo company.CompanyRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {

		companies, err := usecases.CompanySearcher(repo)

		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"code":    http.StatusBadRequest,
				"message": "Something went wrong while processing your request. Invalid filter.",
				"data":    fiber.Map{},
			})
		}

		return c.Status(http.StatusOK).JSON(fiber.Map{
			"items": companies,
		})
	}
}

func CreateCompany(repo company.CompanyRepository, event_broker eventbroker.EventBroker, user_service user.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {

		user_auth_token := c.Get("Authorization", "")

		if user_auth_token == "" {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"code":    http.StatusUnauthorized,
				"message": "Missing Authorization header.",
				"data":    fiber.Map{},
			})
		}

		isValid := user_service.IsUserTokenValid(user_auth_token)

		if !isValid {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"code":    http.StatusUnauthorized,
				"message": "Invalid token.",
				"data":    fiber.Map{},
			})
		}

		var reqBody company.BaseCompany

		if err := c.BodyParser(&reqBody); err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"code":    http.StatusBadRequest,
				"message": "Failed to create record.",
				"data":    fiber.Map{},
			})
		}

		company, err := usecases.CompanyCreator(repo, event_broker, reqBody)

		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"code":    http.StatusBadRequest,
				"message": "Failed to create record.",
				"data":    fiber.Map{},
			})
		}

		return c.Status(http.StatusOK).JSON(company)
	}
}

func FindCompany(repo company.CompanyRepository) fiber.Handler {
	return func(c *fiber.Ctx) error {

		companyId := c.Params("id")

		if companyId == "" {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"code":    http.StatusNotFound,
				"message": "The requested resource wasn't found.",
				"data":    fiber.Map{},
			})
		}

		company, err := usecases.CompanyFinder(repo, companyId)

		if err != nil {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"code":    http.StatusNotFound,
				"message": "The requested resource wasn't found.",
				"data":    fiber.Map{},
			})
		}

		return c.Status(http.StatusOK).JSON(company)
	}
}

func DeleteCompany(repo company.CompanyRepository, event_broker eventbroker.EventBroker, user_service user.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user_auth_token := c.Get("Authorization", "")

		if user_auth_token == "" {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"code":    http.StatusUnauthorized,
				"message": "Missing Authorization header.",
				"data":    fiber.Map{},
			})
		}

		isValid := user_service.IsUserTokenValid(user_auth_token)

		if !isValid {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"code":    http.StatusUnauthorized,
				"message": "Invalid token.",
				"data":    fiber.Map{},
			})
		}

		companyId := c.Params("id")

		if companyId == "" {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"code":    http.StatusNotFound,
				"message": "The requested resource wasn't found.",
				"data":    fiber.Map{},
			})
		}

		err := usecases.CompanyDeleter(repo, event_broker, companyId)

		if err != nil {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"code":    http.StatusNotFound,
				"message": "Failed to delete record.",
				"data":    fiber.Map{},
			})
		}

		return c.Status(http.StatusOK).JSON(nil)
	}
}

func UpdateCompany(repo company.CompanyRepository, event_broker eventbroker.EventBroker, user_service user.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		user_auth_token := c.Get("Authorization", "")

		if user_auth_token == "" {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"code":    http.StatusUnauthorized,
				"message": "Missing Authorization header.",
				"data":    fiber.Map{},
			})
		}

		isValid := user_service.IsUserTokenValid(user_auth_token)

		if !isValid {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"code":    http.StatusUnauthorized,
				"message": "Invalid token.",
				"data":    fiber.Map{},
			})
		}

		companyId := c.Params("id")

		if companyId == "" {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"code":    http.StatusNotFound,
				"message": "The requested resource wasn't found.",
				"data":    fiber.Map{},
			})
		}

		var reqBody company.BaseCompany

		if err := c.BodyParser(&reqBody); err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"code":    http.StatusBadRequest,
				"message": "Failed to update record.",
			})
		}

		err := usecases.CompanyUpdater(repo, event_broker, reqBody, companyId)

		if err != nil {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"code":    http.StatusNotFound,
				"message": "Failed to update record.",
				"data":    fiber.Map{},
			})
		}

		return c.Status(http.StatusOK).JSON(nil)
	}
}
