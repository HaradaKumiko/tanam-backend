# Use golang base image with specified version
FROM golang:1.22.2-alpine3.19 AS build

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the container
COPY go.mod .
COPY go.sum .

# Download dependencies
RUN go mod download

# Copy the rest of the application source code to the container
COPY . .

# Copy the key.json file into the container
COPY utils/google/key.json /app/utils/google/key.json

# Copy the .env file into the container
COPY .env .

# Build the Go app
RUN go build -o main .

# Start a new stage from scratch
FROM alpine:3.19

# Set the working directory inside the container
WORKDIR /root

# Copy the compiled executable from the previous stage
COPY --from=build /app/main .

# Copy the key.json file into the container
COPY --from=build /app/utils/google/key.json /root/utils/google/key.json

# Command to run the executable
CMD ["./main"]