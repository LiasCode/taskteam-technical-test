package router

import (
	"company-microservice/api/container"
	"company-microservice/api/controller"

	"github.com/gofiber/fiber/v2"
)

func CompanyRouter(app *fiber.App) {
	companyApi := app.Group("/company")

	companyApi.Get("/", controller.SearchCompanys(container.ContainerInstance.GetCompanyRepository()))
	companyApi.Get("/:id", controller.FindCompany(container.ContainerInstance.GetCompanyRepository()))
	companyApi.Post("/", controller.CreateCompany(container.ContainerInstance.GetCompanyRepository(), container.ContainerInstance.GetEventBroker(), container.ContainerInstance.GetUserService()))
	companyApi.Patch("/:id", controller.UpdateCompany(container.ContainerInstance.GetCompanyRepository(), container.ContainerInstance.GetEventBroker(), container.ContainerInstance.GetUserService()))
	companyApi.Delete("/:id", controller.DeleteCompany(container.ContainerInstance.GetCompanyRepository(), container.ContainerInstance.GetEventBroker(), container.ContainerInstance.GetUserService()))
}
