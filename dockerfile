# Use the official Golang image as the base image
FROM golang:1.20-alpine

# Set the working directory
WORKDIR /app

# Copy the Go application source code
COPY . .

# Build the Go application
RUN  go mod tidy && go build -o myapp
# will execute any command in a shell inside the container environment

# Expose the port your application will run on
EXPOSE 8080

# Run the application
CMD ["./myapp"]
