FROM golang:1.23.2 AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main cmd/server/main.go

FROM scratch
WORKDIR /app
COPY --from=builder /app/main /app/main
EXPOSE 8000
CMD ["/app/main"]
