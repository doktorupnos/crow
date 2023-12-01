FROM golang:1.21.4 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o backend ./cmd/backend/

FROM scratch
WORKDIR /app
COPY --from=builder /app/backend .
EXPOSE 8000
EXPOSE 8080
CMD ["sh", "-c", "export JWT_SECRET=$(openssl rand -base64 64) && ./backend"]
