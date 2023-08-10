# Use the official Go image as the base image
FROM golang:1.20.6

# Set the working directory in the container
WORKDIR /api

# Copy the application files into the working directory
COPY . /api

# Build the application
RUN go build -o api .

# Expose port 8080
EXPOSE 8080

# Define the entry point for the container
CMD ["./api"]
