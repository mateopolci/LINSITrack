FROM golang:1.25.1-alpine3.22

WORKDIR /app

# Installs Air
RUN go install github.com/air-verse/air@latest

# Copy dependencies first (cache)
COPY go.mod go.sum ./
RUN go mod download

# Copy Air configuration
COPY .air.toml ./

# Copy the rest of the code
COPY . .

# Expose port
EXPOSE 8080

# Start with Air
CMD ["air", "-c", ".air.toml"]