FROM golang:1.21.4 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go test ./... && CGO_ENABLED=0 GOOS=linux go build ./cmd/backend -a -installsuffix cgo -o backend .

FROM scratch
WORKDIR /app
COPY --from=builder /app/backend .
EXPOSE 8080
CMD ["./backend"]
