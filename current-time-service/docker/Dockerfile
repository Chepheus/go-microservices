FROM golang:1.19 as builder

WORKDIR /app

COPY ./current-time-service/go.mod .
COPY ./current-time-service/go.sum .
RUN go mod download && go mod verify

COPY ./current-time-service .

RUN CGO_ENABLED=0 go build -o /app/current-time-service

FROM alpine:3.15.4

WORKDIR /app

COPY --from=builder /app/current-time-service .

EXPOSE 8080

RUN ln -s /app/current-time-service /usr/local/bin/current-time-service

CMD ["current-time-service"]
