FROM golang:alpine as build-env

RUN apk update && apk add git

WORKDIR /src

COPY . .

RUN go mod tidy
RUN go build -o booking-room

FROM alpine
WORKDIR /app
COPY --from=build-env /src/booking-room /app

ENTRYPOINT ./booking-room