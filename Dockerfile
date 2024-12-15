FROM golang:1.23.2-alpine as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o bin/xyz-finance-app cmd/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/bin/xyz-finance-app ./

EXPOSE 4000

CMD ["./xyz-finance-app"]