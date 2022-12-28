# FROM docker.io/library/golang:1.19 as build
# ADD . /src
# WORKDIR /src
# RUN make go/build

# RUN ls /src/bin

# # -------------------------------------------

# # FROM gcr.io/distroless/base
# # FROM alpine
# FROM golang:1.19

# LABEL org.opencontainers.image.source=https://github.com/moabukar/no-hello-bot
# LABEL org.opencontainers.image.description="No Hello Bot for Discord"
# COPY --from=build /src/bin/discord-bot-go /discord-bot-go
# COPY --from=build /src/docs/config.json /etc/config.json
# COPY --from=build /src /app
# USER root:root
# EXPOSE 8888

# # RUN ls /src/bin

# # CMD ["./discord-bot-go", "-c", "/etc/config.json", "migrate"]
# # CMD ["./src/bin/discord-bot-go", "-c", "/etc/config.json", "migrate"]


# ---- WORKING ONE!
FROM golang:1.19


LABEL org.opencontainers.image.source=https://github.com/moabukar/no-hello-bot
LABEL org.opencontainers.image.description="No Hello Bot for Discord"

USER 0:0
EXPOSE 8888

ADD . /src
WORKDIR /src
RUN make go/build
RUN go build /src/cmd/discord-bot-go.go

CMD ["./discord-bot-go"]