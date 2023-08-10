
# Use the official Go image as the base image
FROM golang:1.15

# Set the working directory in the container
WORKDIR /app

# Copy the application files into the working directory
# Error: The error message suggests that the build failed due to an undefined reference.
# To fix this error, we need to make sure that all the necessary dependencies and files are present.
# One common cause for this error is missing Go module files.
# We can use Go modules to manage dependencies and ensure that all required packages are available during.build
# Ensure that all the required dependencies are defined in go.mod and go.sum files.
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy the application source code into the working directory
COPY . .

# Build the application
RUN go build -o main .

# Expose port 8080
EXPOSE 8080

# Define the entry point for the container
CMD ["./main"]
