FROM golang:1.17

WORKDIR /app

COPY go.mod /app/
COPY go.sum /app/
RUN go mod download

COPY ./ /app/
RUN go build ./cmd/main/main.go

EXPOSE 8080

CMD ["./main"]