FROM golang:alpine

RUN mkdir /app

RUN apk add --no-cache ca-certificates

WORKDIR /app

COPY ./binary/authservice .

RUN chmod +x authservice

CMD ["./authservice"]