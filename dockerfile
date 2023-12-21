FROM golang:1.22rc1-bookworm

WORKDIR /app

COPY go.mod .
COPY go.sum .
COPY cmd/main.go cmd/main.go



RUN go mod tidy
RUN go get ./...

RUN GOARCH=amd64 GOOS=linux go build -o bin/main cmd/main.go

ENTRYPOINT ["bin/main"]