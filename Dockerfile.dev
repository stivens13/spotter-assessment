FROM golang:1.23.0-alpine

WORKDIR /app

COPY go.mod ./
RUN go mod download

# COPY app ./app
# COPY data ./data

# RUN CGO_ENABLED=0 GOOS=linux go build -o spotter ./app

EXPOSE 8080

CMD ["./app"]