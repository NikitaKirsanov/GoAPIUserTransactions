FROM golang:alpine

LABEL maintainer="Nikita Kirsanov"

RUN apk update && apk add --no-cache git && apk add --no-cache bash && apk add build-base

RUN mkdir /app
WORKDIR /app

COPY . .
COPY .env .

RUN go get -d -v ./...
RUN go install -v ./...

RUN go get github.com/githubnemo/CompileDaemon
RUN go get -v golang.org/x/tools/gopls

ENTRYPOINT CompileDaemon --build="go build -a -installsuffix cgo -o main ." --command=./main