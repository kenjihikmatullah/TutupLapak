# Start from the latest golang base image
FROM golang:1.23.5 as builder

# Add Maintainer Info
LABEL maintainer="ad1ee"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app for 64-bit ARM architecture
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -a -installsuffix cgo -o main .

######## Start a new stage using a normal Linux distribution ########
FROM debian:12-slim

# Install certificates
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*

# Set the working directory
WORKDIR /root/

# Copy the Pre-built binary file from the previous stage
COPY --from=builder /app/main .
COPY --from=builder /app/.env .

# Tentukan argumen untuk port
ARG PORT

# Set port sebagai environment variable
ENV PORT=${PORT}

# Tentukan port yang akan diekspos
EXPOSE ${PORT}

# Command to run the executable
CMD ["./main"]
