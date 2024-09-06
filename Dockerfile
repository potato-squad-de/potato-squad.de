FROM alpine:latest as runner
WORKDIR /root/
COPY site /usr/bin/
ENTRYPOINT ["site", "run"]

ENV PORT=80
ENV DISCORD_INVITE_URL=

EXPOSE ${PORT}
