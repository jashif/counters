# Start from a lightweight image with Go installed
FROM golang:1.18-alpine as builder

WORKDIR /app

# Copy the go mod and sum files to download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd

# Use a scratch (empty) image to run the binary
FROM scratch

WORKDIR /root/

# Copy the pre-built binary file from the previous stage
COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]
