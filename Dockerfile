FROM golang:1.16.4-alpine

WORKDIR /src

COPY go.mod go.sum ./

RUN go mod download

COPY . .

ENV PORT=${PORT}

RUN go get -u github.com/cosmtrek/air

RUN touch .air.toml

WORKDIR /src/cmd

CMD ["air"]  