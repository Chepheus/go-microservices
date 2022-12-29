FROM golang:1.19 as builder

WORKDIR /app

#COPY go.mod go.sum ./
#RUN go mod download && go mod verify

COPY ./metrics-service .
RUN CGO_ENABLED=0 go build -o /app/metrics-service

FROM alpine:3.15.4

WORKDIR /app

COPY --from=builder /app/metrics-service .

RUN ln -s /app/metrics-service /usr/local/bin/metrics-service

CMD ["metrics-service"]
