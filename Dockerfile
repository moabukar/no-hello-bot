FROM docker.io/library/golang:1.19 as build
ADD . /src
WORKDIR /src
RUN make go/build

# -------------------------------------------

FROM gcr.io/distroless/base
# LABEL org.opencontainers.image.licenses=AGPLv3
LABEL org.opencontainers.image.source=https://github.com/moabukar/no-hello-bot
LABEL org.opencontainers.image.description="No Hello Bot for Discord"
COPY --from=build /src/cmd/discord-bot-go /discord-bot-go
COPY --from=build /src/docs/config.json /etc/config.json
USER 10000:10000
EXPOSE 8888
# ENTRYPOINT [ "/", "-c", "/etc/config.json", "--migrate" ]

ENTRYPOINT [ "/discord-bot-go"]

# ENTRYPOINT [ "/discord-bot-go", "-c", "/etc/config.json", "--migrate" ]