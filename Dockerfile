# syntax=docker/dockerfile-upstream:master-labs
FROM golang:1.12.4-alpine3.9 AS builder

WORKDIR /go/src/lupusmic.org/rip
RUN apk add \
    --no-cache \
    --update \
    alpine-sdk

COPY ./src/ /go/src/lupusmic.org/rip
RUN go get -t ./...
RUN go build

FROM alpine:3.9
WORKDIR /go/src/lupusmic.org/rip
ENTRYPOINT [ "/entrypoint" ]
CMD [ "rip", "--config", "/etc/lupusmic.org/rip.json" ]

RUN apk add \
    --no-cache \
    --update \
    ca-certificates openssl

COPY ./entrypoint /entrypoint
COPY --from=builder /go/bin/rip /usr/local/bin/rip
COPY \
    ./src/graphql/schema.graphql \
    /go/src/lupusmic.org/rip/graphql/schema.graphql
COPY \
    config-docker.json \
    /etc/lupusmic.org/rip.json
