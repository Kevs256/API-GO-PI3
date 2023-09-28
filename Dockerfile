#build stage
FROM golang:1.21.1-alpine3.18 AS builder
RUN apk add --no-cache git

WORKDIR /app

COPY ["go.mod", "go.sum", "./"]
RUN go mod download

COPY . .

RUN go build -o app -v .

#final stage
FROM alpine:3.18
LABEL Name=apigoauth Version=0.0.1

RUN apk update
RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /app .
ENTRYPOINT ["./app"]
