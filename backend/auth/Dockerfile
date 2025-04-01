FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY auth/go.mod auth/go.sum ./
COPY shared/ ./shared/
RUN go mod edit -replace asl-great-mind/shared=./shared
RUN go mod download
COPY auth/ .
RUN go build -o auth cmd/api/main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/auth .
EXPOSE 8081
CMD ["./auth"]