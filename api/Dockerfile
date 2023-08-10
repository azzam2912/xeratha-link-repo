FROM golang:1.20.6

WORKDIR /api

COPY . /api

RUN go build -o main .

EXPOSE 8080

CMD ["./api/main"]
