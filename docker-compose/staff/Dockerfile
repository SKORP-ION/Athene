FROM alpine:3.14

COPY staff /
RUN chmod +x staff

RUN apk update
RUN apk add tzdata

CMD ["./staff"]
