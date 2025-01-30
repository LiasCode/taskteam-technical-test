FROM golang:1.23

WORKDIR /app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

# Install Make utility
RUN apt-get update && apt-get install -y make

# Build the Go app
RUN make build

# Command to run the executable
CMD ["make", "start"]
