FROM alpine:3.14

COPY ldap /
RUN chmod +x ldap

RUN apk update
RUN apk add tzdata

CMD ["./ldap"]
