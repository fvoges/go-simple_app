FROM golang:1.17-alpine AS build

LABEL name="simple_app"
LABEL org.opencontainers.image.title="simple_app"
LABEL org.opencontainers.image.vendor="Federico Voges"
LABEL org.opencontainers.image.authors="fvoges@gmail.com"
LABEL org.opencontainers.image.version="0.1.1"
LABEL org.opencontainers.image.description="Simple web app that prints the content of a JSON file"

WORKDIR /src/
COPY main.go go.* /src/
RUN CGO_ENABLED=0 go build -o /bin/demo

# FROM scratch
FROM alpine:3.14
COPY --from=build /bin/demo /bin/demo
COPY ./config.json /config.json

EXPOSE 8888
ENV JSON_FILE=/config.json
ENTRYPOINT ["/bin/demo"]


