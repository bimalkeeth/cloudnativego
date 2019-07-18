FROM alpine:3.5

MAINTAINER Bimal Kaluarachchi

COPY ./cloudnativego /app/cloudnativego

RUN chmod +x /app/cloudnativego

EXPOSE 8080

ENTRYPOINT /app/cloudnativego


