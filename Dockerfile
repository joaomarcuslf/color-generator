FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

ENV PORT 5003

EXPOSE 5003

CMD ["./main"]
