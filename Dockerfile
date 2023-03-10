FROM golang:1.20.0-alpine3.17

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .

RUN go mod tidy 

RUN go build -o binary

ENTRYPOINT ["/app/binary"]