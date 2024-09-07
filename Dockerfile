ARG GO_VERSION=1.20

FROM golang:$GO_VERSION-alpine as deps

WORKDIR /app

COPY ./go.mod ./go.sum ./

RUN go mod download -x

###########################################################

FROM deps as builder-main
WORKDIR /app
COPY . .
RUN pwd && go build -o /app/main ./app/cmd/main.go
#  ##################################################

FROM alpine:latest

WORKDIR /app

COPY --from=builder-main /app/main /app/main

COPY ./app/config/config.yaml /app/config.yaml

EXPOSE 8080

CMD ["/app/main"]
