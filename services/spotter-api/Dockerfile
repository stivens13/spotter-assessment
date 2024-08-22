# staging build
FROM golang:1.23.0-alpine AS builder
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY services/spotter-api ./services/spotter-api
COPY services/youtube-client ./services/youtube-client
COPY config ./config
COPY helper ./helper
COPY models ./models

RUN CGO_ENABLED=0 GOOS=linux go build -o spotter-api ./services/spotter-api

# production build
FROM alpine:3.20.2

WORKDIR /app

COPY --from=builder /app/spotter-api ./spotter-api

RUN apk --no-cache add ca-certificates tzdata

EXPOSE 8080

CMD ["./spotter-api"]