FROM golang:1.17.1-alpine3.14 AS build

RUN go get -u github.com/cosmtrek/air

ENV GO111MODULE=on
WORKDIR /go/src/nsym-m/chat-api
COPY ./src/go.mod ./src/go.sum ./
RUN go mod download

COPY ./src/. ./
RUN go build

FROM alpine:3.13
COPY --from=build /go/src/nsym-m/go-rest /usr/local/bin/

EXPOSE 8080
CMD ["air", "-c", ".air.toml"]