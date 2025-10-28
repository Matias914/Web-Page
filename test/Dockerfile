FROM alpine:latest

RUN apk add --no-cache bash curl jq

COPY ./test /app/test

RUN chmod +x /app/test/*.sh

WORKDIR /app