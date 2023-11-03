FROM golang:1.18-alpine as builder

RUN apk add alpine-sdk

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /go-crud

EXPOSE 5000

CMD ["/go-crud"]
