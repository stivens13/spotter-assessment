# etl service Dockerfile
# staging build
FROM golang:1.23.0-alpine AS builder
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY services/etl ./services/etl
COPY services/youtube-client ./services/youtube-client
COPY tools/generator ./tools/generator
COPY config ./config
COPY helper ./helper
COPY models ./models

RUN CGO_ENABLED=0 GOOS=linux go build -o etl ./services/etl

# production build
FROM alpine:3.20.2

WORKDIR /app

COPY --from=builder /app/etl ./etl

RUN apk --no-cache add curl ca-certificates tzdata

EXPOSE 8080

CMD ["./etl"]