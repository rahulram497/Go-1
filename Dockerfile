FROM golang:1.24 AS builder

WORKDIR /

COPY . .

RUN go build -o go-1

EXPOSE 5000

CMD ["./go-1"]
