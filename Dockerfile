FROM golang:1.16.4-alpine as build

ENV GOOS=linux \
    GOARCH=amd64

WORKDIR /app

COPY go.mod ./

RUN set -ex; \
    apk update; \
    apk add --no-cache \
    git

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o goapp main.go

##########################################

FROM alpine:3.13.1 as release

WORKDIR /app

RUN apk add --no-cache ca-certificates

COPY --from=build /app/goapp ./goapp

EXPOSE 3000

CMD ["./goapp"]
