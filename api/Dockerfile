FROM golang:1.21.4 as builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go test ./... && CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o api .

FROM scratch
WORKDIR /app
COPY --from=builder /app/api .
EXPOSE 8080
CMD ["./api"]
