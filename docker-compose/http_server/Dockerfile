FROM alpine:latest

COPY http_server /http_server/http_server
RUN chmod +x /http_server/http_server

COPY templates /http_server/templates
COPY static /http_server/static
COPY favicon.ico /http_server/favicon.ico
COPY cert.pem /http_server/cert.pem
COPY privkey.pem /http_server/privkey.pem

WORKDIR /http_server

RUN apk update && apk upgrade
RUN apk add --no-cache --update openssl ca-certificates curl tzdata

CMD ["/http_server/http_server"]
