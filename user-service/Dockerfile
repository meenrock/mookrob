# Use the official Go image as a parent image
FROM golang:1.21 AS build

# Set the working directory inside the container
WORKDIR /app

# Copy the Go application source code into the container
COPY . .

WORKDIR /app/src

# Build the Go application
RUN go build -o myapp ./main

# Start a new stage to create a minimal image
FROM scratch

# Copy the compiled binary from the previous stage
COPY --from=build /app/src/myapp /myapp

# Expose the port the application will run on (if needed)
EXPOSE 8080

# Define the command to run your application
CMD ["/myapp"]