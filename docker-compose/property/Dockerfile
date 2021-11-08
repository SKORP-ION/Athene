FROM alpine:3.14

COPY property /
RUN chmod +x property

RUN apk update
RUN apk add tzdata

CMD ["./property"]
