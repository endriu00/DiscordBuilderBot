FROM golang:1.16-alpine AS build

WORKDIR /bot

COPY go.* ./
RUN go mod download

COPY cmd cmd
COPY service service 

RUN mkdir /app

WORKDIR /bot/cmd/
RUN go build -o /app/ ./

FROM alpine:latest

WORKDIR /app/

COPY --from=build /app/* ./

ENTRYPOINT [ "/app/cmd" ]
