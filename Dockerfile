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
FROM golang:alpine3.17 AS builder

ENV TOKEN=$TOKEN

LABEL org.opencontainers.image.source=https://github.com/moabukar/no-hello-bot
LABEL org.opencontainers.image.description="No Hello Bot for Discord"
ENV GOARCH=amd64


USER 0:0
EXPOSE 8888

ADD . /src
WORKDIR /src
RUN apk update && apk add --no-cache make
RUN go build -o discord-bot-go /src/cmd/discord-bot-go.go
# RUN make go/build
RUN go build /src/cmd/discord-bot-go.go
# RUN echo '{"key":"value"}' > /src/docs/config.json

COPY .env /app/.env
ENV TOKEN_FILEPATH=/app/.env


# CMD ["./discord-bot-go"]

# FROM gcr.io/distroless/base

FROM gcr.io/distroless/static

# FROM scratch


COPY --from=builder /src/discord-bot-go /app/discord-bot-go
COPY --from=builder /src/docs/config.json ./docs/config.json

# USER nobody:nobody
## Use IDs instead >> nobody=65534
USER 65534:65534

ENTRYPOINT ["/app/discord-bot-go"]