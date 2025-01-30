package main

import (
	"company-microservice/cmd/server"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	server.InitServer()
}
