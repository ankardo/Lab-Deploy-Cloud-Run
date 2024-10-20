FROM golang:1.23.2 AS golang
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

FROM golang AS builder
WORKDIR /build
COPY . .
RUN  go build -o main cmd/server/main.go && \
  chmod +x main

FROM scratch
WORKDIR /app
COPY --from=builder /build/main /app/main
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY .env ./
EXPOSE 8000
CMD ["/app/main"]
