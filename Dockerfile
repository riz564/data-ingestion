# Start from the official Go image for building
FROM golang:1.23 as builder

WORKDIR /app

# Copy go mod and sum files
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy the source code
COPY . .

# Build the Go app
RUN go build -o data-ingestion

# Use a minimal image for running
FROM gcr.io/distroless/base-debian12

WORKDIR /app

# Copy the built binary
COPY --from=builder /app/data-ingestion .

# Copy credentials if needed (optional, see docker-compose for mounting secrets)

# Set environment variables (can also be set in docker-compose)
# ENV GCP_PROJECT_ID=your-project-id
# ENV GOOGLE_APPLICATION_CREDENTIALS=/app/service-account.json

CMD ["./data-ingestion"]