FROM golang:1.23.1

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o main .

CMD ["./main"]
