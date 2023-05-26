# Development Image
FROM golang:1.19-alpine AS release

ENV ENV development

# Install git, which is required for 'go get'
RUN apk update && apk add --no-cache git

# Needed for VS Code Remote Development
RUN apk add musl libgcc libstdc++ curl wget bash ca-certificates gcompat openssh-client

RUN go install github.com/cosmtrek/air@latest

# Set the Current Working Directory inside the container
WORKDIR /go/src/fiber-book-app

# Copy all files from parent directory to the current working directory
COPY . .

# Downloads all the dependencies in advance (could be left out, but it's more clear this way)
RUN go mod download

EXPOSE 9000

# chmod "dev.sh" so it can be executed
RUN chmod +x dev.sh

# Run the executable
CMD ["ash", "./dev.sh"]


# Production Image
# Building the binary of the App
FROM golang:1.19-alpine AS build

ENV ENV production

# Install git, which is required for 'go get'
RUN apk update && apk add --no-cache git

# Set the Current Working Directory inside the container
WORKDIR /go/src/fiber-book-app

# Copy all the Code and stuff to compile everything
COPY . .

# Downloads all the dependencies in advance (could be left out, but it's more clear this way)
RUN go mod download

# Builds the application as a staticly linked one, to allow it to run on alpine
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o app .

EXPOSE 9000

# Run the executable
CMD ["./app"]