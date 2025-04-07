FROM golang:1.24.2-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build \
  -tags netgo \
  -ldflags="-w -s -extldflags '-static'" \
  -trimpath \
  -o /api

FROM scratch

WORKDIR /app

# Copy the built api binary
COPY --from=builder /api .

# Copy the database migration files
COPY --from=builder /app/sql/schema sql/schema

EXPOSE 8000

ENTRYPOINT ["/app/api"]

