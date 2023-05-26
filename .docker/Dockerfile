# Development Image
FROM golang:1.19 AS release

ENV ENV development

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
CMD ["bash", "./dev.sh"]

