# syntax=docker/dockerfile:1

FROM golang:1.19.0-alpine as build
WORKDIR /src
COPY go.mod .
COPY go.sum .
RUN go mod download && go mod verify

COPY . .
RUN go build ./cmd/main.go

FROM alpine:latest
COPY --from=build /src/main /history
COPY --from=build /src/config/config.json /config/config.json

# api port
EXPOSE 80

# rpc port
EXPOSE 8000

CMD ["./history", "-race"]