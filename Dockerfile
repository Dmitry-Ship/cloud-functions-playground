FROM golang:1.16.4-alpine

WORKDIR /src

COPY go.mod go.sum ./

RUN go mod download

COPY . .

ENV PORT=${PORT}

RUN go get github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon --build="go build -v -o tmp/bin ./cmd/main.go" --command=./tmp/bin

