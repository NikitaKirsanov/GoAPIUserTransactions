FROM golang

LABEL maintainer="Nikita Kirsanov"

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./
COPY .env .

RUN go build -o /main

EXPOSE 8080

CMD [ "/main" ]
