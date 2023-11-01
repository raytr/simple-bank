# syntax=docker/dockerfile:1
FROM golang:1.19-alpine AS builder

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod go.sum ./
RUN go mod download

# Copy everything from this project into the filesystem of the container.
COPY . .

# Compile the binary exe for our app.
RUN CGO_ENABLED=0 GOOS=linux go build -o /simple-bank

EXPOSE 8081

# Start the application.
CMD ["/simple-bank"]