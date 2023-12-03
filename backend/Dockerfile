FROM golang:1.21.4 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o backend ./cmd/backend/

FROM alpine:3.18.2
WORKDIR /app
COPY --from=builder /app/backend .
RUN apk add --no-cache openssl
EXPOSE 8000
EXPOSE 8080
CMD ["sh", "-c", "export JWT_SECRET=$(openssl rand -base64 64) && ./backend"]
