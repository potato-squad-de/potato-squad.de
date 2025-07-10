FROM docker.io/hugomods/hugo:node AS builder
WORKDIR /src
COPY . .
RUN npm i && hugo

FROM caddy:latest AS runner
COPY --from=builder /src/public /srv
ENTRYPOINT ["caddy", "file-server" , "--root", "/srv"]