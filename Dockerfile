FROM golang:1.17-alpine AS builder

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o welcome-app cmd/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/welcome-app .

EXPOSE 5000

CMD ["./welcome-app"]