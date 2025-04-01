FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY core/go.mod core/go.sum ./
COPY shared/ ./shared/
RUN go mod edit -replace asl-great-mind/shared=./shared
RUN go mod download
COPY core/ .
RUN go build -o core cmd/api/main.go

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/core .
EXPOSE 8080
CMD ["./core"]