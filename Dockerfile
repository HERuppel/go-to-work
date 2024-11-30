FROM golang:1.23.3-alpine

RUN apk update && apk add --no-cache make curl bash

RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar -xz -C /usr/local/bin

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o main cmd/api/main.go

COPY entrypoint.sh /entrypoint.sh

RUN chmod +x /entrypoint.sh

EXPOSE 3333

ENTRYPOINT ["/entrypoint.sh"]


CMD ["./main"]