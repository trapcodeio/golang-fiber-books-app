FROM golang:1.19 AS release

ENV ENV development

RUN go install github.com/cosmtrek/air@latest

# Set the Current Working Directory inside the container
WORKDIR /go/src/fiber-book-app

# Copy all files from parent directory to the current working directory
COPY ../. /go/src/fiber-book-app

# list all files in the current directory
RUN pwd
RUN ls -la

# Downloads all the dependencies in advance (could be left out, but it's more clear this way)
RUN go mod download

EXPOSE 9000

# chmod "dev.sh" so it can be executed
RUN chmod +x dev.sh

# Run the executable
CMD ["bash", "./dev.sh"]


# Production Image
# Building the binary of the App
FROM golang:1.19 AS build

ENV ENV production

# Set the Current Working Directory inside the container
WORKDIR /go/src/fiber-book-app

# Copy all the Code and stuff to compile everything
COPY ../. /go/src/fiber-book-app

# Downloads all the dependencies in advance (could be left out, but it's more clear this way)
RUN go mod download


# Builds the application as a staticly linked one, to allow it to run on alpine
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o app .

EXPOSE 9000

# Run the executable
CMD ["./app"]
