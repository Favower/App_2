# Dockerfile
FROM golang:1.20-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
COPY config/*.go ./config/
COPY handlers/*.go ./handlers/
COPY models/*.go ./models/
COPY services/*.go ./services/

RUN go build -o /go-microservice

EXPOSE 8080

CMD [ "/go-microservice" ]
