FROM alpine:latest

RUN apk update && \
    apk add --no-cache \
        php \
        php-fpm \
        php-mysqli \
        php-json \
        php-openssl \
        php-mbstring \
        php-xml \
        php-ctype \
        php-session \
        go

WORKDIR /src

RUN php -v && go version

CMD ["/bin/sh"]
