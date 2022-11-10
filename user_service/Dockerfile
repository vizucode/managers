FROM golang:alpine

RUN mkdir /app

RUN apk add --no-cache ca-certificates

WORKDIR /app

COPY ./binary/userservice .

RUN chmod +x userservice

CMD ["./userservice"]