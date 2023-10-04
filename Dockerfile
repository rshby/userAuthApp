FROM golang:1.21.1-alpine as builder

WORKDIR /app

COPY ./ ./

RUN go mod tidy
RUN go build -o ./bin/appauth ./main.go

FROM alpine:3

WORKDIR /app

COPY --from=builder /app/.env ./
COPY --from=builder /app/bin/appauth ./

EXPOSE 5005
CMD ./appauth