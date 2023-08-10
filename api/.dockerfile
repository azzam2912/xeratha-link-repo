FROM golang:1.20.6

WORKDIR /api

COPY . /api

RUN go build -o cmd/main .

EXPOSE 8080

CMD ["./cmd/main"]