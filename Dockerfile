FROM golang:alpine as builder

LABEL maintainer="Egor Miloserdov"

RUN apk update && apk add --no-cache git

WORKDIR /atlant-test
COPY . .
COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o main .


### making a building container
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /atlant-test/main .
COPY --from=builder /atlant-test/.env .
#ENV DB_HOST="host.docker.internal"

EXPOSE 8080

ENTRYPOINT ./main
