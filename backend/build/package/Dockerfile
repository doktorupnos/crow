FROM golang:1.23-alpine as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
ENV CGO_ENABLED=0
ENV GOOS=linux
RUN go build -a -installsuffix cgo -ldflags="-w -s" -o backend cmd/backend/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/backend .
EXPOSE 8000
CMD ["./backend"]
