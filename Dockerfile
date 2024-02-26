# Stage 1: Build the application
FROM golang:1.17 AS builder

WORKDIR /app

# Copy go.mod and go.sum files to download dependencies
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the entire project to the working directory
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Stage 2: Create a minimal runtime image
FROM alpine:latest  

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the executable from the builder stage
COPY --from=builder /main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
