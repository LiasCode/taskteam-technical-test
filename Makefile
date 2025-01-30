dev:
	air 

build:
	go build -o ./dist/company-microservice

start: build
	./dist/company-microservice

test: 
	go test -v company-microservice/pkg/Company
	# go test -v company-microservice/pkg/Company/UseCases
	go test -v company-microservice/internals
	go test -v company-microservice/api

docker-build:
	docker build -t company-microservice .

docker-run:
	docker compose up -d

docker-stop:
	docker compose down

docker-remove:
	docker rm company-microservice-container

docker-clean:
	docker rmi company-microservice

help:
	@echo "Available commands:"
	@echo "  dev            - Start the development server"
	@echo "  build          - Build the project into the dist/ folder"
	@echo "  start          - Build and start the production server"
	@echo "  docker-build   - Build the Docker image for the application"
	@echo "  docker-run     - Run the application using Docker Compose"
	@echo "  docker-stop    - Stop all running containers with Docker Compose"
	@echo "  docker-remove  - Remove the stopped Docker container"
	@echo "  docker-clean   - Remove the Docker image"
