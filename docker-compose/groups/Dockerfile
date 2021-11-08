FROM alpine:3.14

COPY groups /
RUN chmod +x groups

RUN apk update
RUN apk add tzdata

CMD ["./groups"]
