# Use the official Go image as the base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy your Go application files to the container
COPY . .

# Build your Go application
RUN go build -o main

# Expose the port your Go API will listen on
EXPOSE 8080

# Command to run your Go API
CMD ["./main"]
