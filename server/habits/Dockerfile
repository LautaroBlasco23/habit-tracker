FROM golang:1.22

WORKDIR /app

COPY . .

RUN go build

EXPOSE 8080

CMD ["./habitservice"]
