package api

import (
	"company-microservice/api/container"
	"os"
	"testing"
)

func TestIntegration(t *testing.T) {
	container.ContainerInstance.IsTesting = true
	os.Setenv("PORT", "8080")

}
