FROM golang:1.24.2-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build \
  -tags netgo \
  -ldflags="-w -s -extldflags '-static'" \
  -trimpath \
  -o /api ./cmd/api/main.go

FROM scratch

WORKDIR /app
COPY --from=builder /api .

EXPOSE 8000

ENTRYPOINT ["/app/api"]

