FROM golang:alpine

LABEL maintainer="Nikita Kirsanov"

RUN apk update && apk add --no-cache git && apk add --no-cache bash && apk add build-base

RUN mkdir /app
WORKDIR /app

COPY . .
COPY .env .

RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 8080
CMD [ "/app/main" ]