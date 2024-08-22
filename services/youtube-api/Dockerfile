# youtube-api service Dockerfile
# staging build
FROM golang:1.23.0-alpine AS builder
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY services/youtube-api ./services/youtube-api
COPY tools/generator ./tools/generator
COPY helper ./helper

RUN CGO_ENABLED=0 GOOS=linux go build -o youtube-api ./services/youtube-api

# production build
FROM alpine:3.20.2

WORKDIR /app

COPY --from=builder /app/youtube-api ./youtube-api

RUN apk --no-cache add curl ca-certificates tzdata

EXPOSE 9000

CMD ["./youtube-api"]