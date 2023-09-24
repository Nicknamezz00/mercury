# protobuf
FROM golang:1.21-alpine AS protobuf

WORKDIR /protobuf-gen
COPY . .
RUN GO111MODULE=on GOBIN=/usr/local/bin go install github.com/bufbuild/buf/cmd/buf@v1.26.1

WORKDIR /protobuf-gen/proto
RUN buf generate

# build frontend

# build backend

FROM alpine:latest AS monolithic
LABEL authors="aaron"

WORKDIR /usr/local/mercury

RUN apk add --no-cache tzdata
ENV TZ="UTC"

COPY --from=backend /backend/mercury /usr/local/mercury/

EXPOSE 3333

RUN mkdir -p /var/opt/mercury
VOLUME /var/opt/mercury

ENV MERCURY_MODE="production"
ENV MERCURY_PORT="3333"

ENTRYPOINT ["./mercury"]