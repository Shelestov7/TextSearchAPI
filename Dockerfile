# Start from a base image that includes Go
FROM golang:1.19-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go application files to the container
COPY go.mod go.sum ./

# download dependencies
RUN go mod download

# Copy all files in container
COPY . .

# Bind port
EXPOSE 8080

# Build the Go application
RUN go build -o /app/main ./cmd

# Start the Go application
CMD ["/app/main"]