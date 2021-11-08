FROM alpine:3.14

COPY history /
RUN chmod +x history

RUN apk update
RUN apk add tzdata

CMD ["./history"]
