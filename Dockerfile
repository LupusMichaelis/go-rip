FROM golang:1.12.4-alpine3.9 AS builder

RUN apk add \
    --no-cache \
    --update \
    alpine-sdk

COPY . /go/src/lupusmic.org/rip
WORKDIR /go/src/lupusmic.org/rip
RUN go get -t ./...
RUN go build
RUN openssl \
    req \
    -subj '/CN=lupusmic.org/O=None/C=FR' \
    -newkey rsa:2048 \
    -nodes \
    -keyout server.key \
    -x509 \
    -days 365 \
    -out server.crt

RUN chmod 400 server.key server.crt

FROM alpine:3.9

RUN apk add \
    --no-cache \
    --update \
    ca-certificates

COPY --from=builder /go/bin/rip /usr/local/bin/rip
COPY --from=builder \
    /go/src/lupusmic.org/rip/graphql/schema.graphql \
    /go/src/lupusmic.org/rip/graphql/schema.graphql
COPY --from=builder \
    /go/src/lupusmic.org/rip/server.crt \
    /go/src/lupusmic.org/rip/server.crt 
COPY --from=builder \
    /go/src/lupusmic.org/rip/server.key \
    /go/src/lupusmic.org/rip/server.key 
COPY --from=builder \
    /go/src/lupusmic.org/rip/config-docker.json \
    /etc/lupusmic.org/rip.json

WORKDIR /go/src/lupusmic.org/rip
CMD [ "rip", "--config", "/etc/lupusmic.org/rip.json" ]
