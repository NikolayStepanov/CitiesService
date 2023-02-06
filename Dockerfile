FROM golang

WORKDIR /app

COPY go.mod .

RUN go mod download

COPY . .

RUN GOOS=linux go build -o ./.bin/app ./cmd/app/main.go

ENV PORT 8080

VOLUME [ "/app/data" ]

EXPOSE $PORT

CMD ["./.bin/app"]

