FROM golang:1.25.1-alpine3.22

WORKDIR /app

# Copy dependencies first (cache)
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the code
COPY . .

# Expose port
EXPOSE 8080

# Start the application directly with Go
CMD ["go", "run", "main.go"]