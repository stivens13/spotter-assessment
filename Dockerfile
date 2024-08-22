# staging build
FROM golang:1.23.0-alpine AS builder
WORKDIR /workdir

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY app ./app

RUN CGO_ENABLED=0 GOOS=linux go build -o spotter-api ./app

# production build
FROM alpine:3.20.2

WORKDIR /app

COPY --from=builder /workdir/spotter-api ./spotter-api

RUN apk --no-cache add ca-certificates tzdata

EXPOSE 8080

CMD ["./spotter-api"]