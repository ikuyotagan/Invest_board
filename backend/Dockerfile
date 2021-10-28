FROM golang:1.17-alpine

WORKDIR /app

COPY . ./

RUN apk update && apk upgrade && apk add --no-cache git

RUN go mod download

RUN go build -o /exe ./cmd/apiserver/main.go

ENTRYPOINT [ "/exe" ]

