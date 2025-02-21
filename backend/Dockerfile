FROM golang:latest

WORKDIR /go/src/app

COPY . .

RUN go mod download

RUN go build -o backend ./cmd/labs/main.go

CMD ["./backend"]