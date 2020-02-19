FROM golang:1.13

WORKDIR /cmd/simple_chat

COPY go.mod ./
RUN go mod download 
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

CMD ["./main"]