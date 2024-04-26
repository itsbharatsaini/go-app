FROM golang:1.22-alpine

EXPOSE 8080

WORKDIR /app

COPY . .

RUN go build -o myapp .

CMD ["./myapp"]