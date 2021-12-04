FROM golang:1.17-alpine3.14 

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o prog

EXPOSE 8000

CMD ["./prog"]
