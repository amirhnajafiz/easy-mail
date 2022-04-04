FROM golang:1.17-alpine as builder

WORKDIR /app/src/

RUN apk add --no-cache git

COPY . .

WORKDIR /app/src/cmd

RUN go build -o ./runner

FROM alpine:latest

WORKDIR /app/

COPY --from=builder ./runner .

ENTRYPOINT ["./runner"]