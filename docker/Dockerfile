FROM golang:alpine

WORKDIR /app

COPY ./app .

RUN go build -o main .

EXPOSE 5000

CMD ["./main"]
