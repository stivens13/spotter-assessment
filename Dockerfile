FROM golang:1.23.0-alpine

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY app ./app
COPY data ./data

EXPOSE 8080

RUN CGO_ENABLED=0 go install -ldflags "-s -w -extldflags '-static'" github.com/go-delve/delve/cmd/dlv@latest

RUN CGO_ENABLED=0 go build -gcflags "all=-N -l" -o app-debugger ./app
CMD [ "/go/bin/dlv", "--listen=:4000", "--headless=true", "--log=true", "--accept-multiclient", "--api-version=2", "exec", "/app/app-debugger" ]