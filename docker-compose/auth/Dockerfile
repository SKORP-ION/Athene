FROM alpine:3.14

COPY auth /
RUN chmod +x auth
RUN apk update
RUN apk add tzdata

CMD ["./auth"]
