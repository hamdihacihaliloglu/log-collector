FROM golang:1.22

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
COPY wait-for-it.sh /wait-for-it.sh

WORKDIR /app/cmd/server

RUN go build -o main .

CMD ["./main"]
