FROM golang:alpine

RUN apk update && apk add --no-cache git

WORKDIR /opt/app

COPY . .

RUN go mod tidy

RUN go build -o bin/app

ENTRYPOINT ["bin/app"]