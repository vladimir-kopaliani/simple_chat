FROM golang:1.13

WORKDIR /opt

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/simple_chat/main.go

CMD ["./main"]
