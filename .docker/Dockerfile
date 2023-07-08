# Development Image
FROM golang:1.19-alpine AS dev

ENV ENV development

# Needed for VS Code Remote Development
RUN apk add musl libgcc libstdc++ git curl wget bash ca-certificates gcompat openssh-client

# Set the Current Working Directory inside the container
WORKDIR /usr/src/app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy all files from parent directory to the current working directory
COPY . .

# Install air for hot reloading
# binary will be $(go env GOPATH)/bin/air
RUN curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

# Build the code to avoid air building the code at startup
RUN go build .

EXPOSE 9000

# Run the executable
CMD ["air"]


# Production Image
# Building the binary of the App
FROM golang:1.19-alpine AS prod

ENV ENV production

# Install git, which is required for 'go get'
RUN apk update && apk add --no-cache git

# Set the Current Working Directory inside the container
WORKDIR /usr/src/app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy all files from parent directory to the current working directory
COPY . .

# Builds the application as a staticly linked one, to allow it to run on alpine
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o app .

EXPOSE 9000

# Run the executable
CMD ["./app"]