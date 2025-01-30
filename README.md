# Companies Microservice

## Requirements

- Go 1.20
- Docker
- Docker Compose
- Air
- Make

## Usage

### `make dev`

- Starts the development server using `air`
- This command is used to run the application in development mode, where code changes are automatically detected and the server is restarted.

### `make build`

- Builds the Go application using `go build`
- This command compiles the Go code and creates an executable file in the `dist` directory.

### `make start`

- Builds the Go application using `go build`
- Starts the production server using the executable file in the `dist` directory
- This command is used to run the application in production mode.

### `make test`

- Runs the unit tests for the application using `go test`
- This command executes the test cases written for the application and reports any failures or errors.

### `make docker-build`

- Builds the Docker image for the application using `docker build`
- This command creates a Docker image with the application code and dependencies.

### `make docker-run`

- Runs the Docker container using `docker compose up`
- This command starts the Docker container and makes the application available at `http://localhost:3030`.

### `make docker-stop`

- Stops the Docker container using `docker compose down`
- This command stops the Docker container and removes any resources associated with it.

### `make docker-remove`

- Removes the Docker container using `docker rm`
- This command deletes the Docker container and any associated resources.

### `make docker-clean`

- Removes the Docker image using `docker rmi`
- This command deletes the Docker image and any associated resources.

### `make help`

- Displays a list of available `make` scripts and their descriptions
- This command provides a quick reference for the available `make` scripts and their purposes.

---

### Project Folder Structure

- `api`: Folder containing the implementation of services and infrastructure
- `cmd`: Folder containing the application commands and server startup
- `internals`: Folder containing utility and reusable files
- `pkg`: Folder containing the domain and business logic
  - `Company`
    - `model`
    - `repository`
    - `usecases`
  - `EventBroker`
  - `User`
- `Makefile`: File containing instructions to compile and run the application
- `docs`: Contains the project documentation
