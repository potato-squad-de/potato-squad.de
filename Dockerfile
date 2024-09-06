FROM alpine:latest as runner
WORKDIR /root/
COPY site /usr/bin/
ENTRYPOINT site
