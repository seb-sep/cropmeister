
# Use the official Golang image as the base image
FROM golang:1.16

# Set the working directory inside the container
WORKDIR /app

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o main .
RUN chmod +x main
RUN ./main

# Set the entry point for the container
ENTRYPOINT ["./main"]
