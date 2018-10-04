# Build Geth in a stock Go builder container
FROM golang:1.9-alpine as builder

RUN apk add --no-cache make gcc musl-dev linux-headers

ADD . /cryptopepe-reader
RUN cd /cryptopepe-reader && build/env.sh go build -v -o build/cryptopepe-reader .

# Pull Geth into a second stage deploy alpine container
FROM alpine:latest

RUN apk add --no-cache ca-certificates
COPY --from=builder /cryptopepe-reader/build/cryptopepe-reader /usr/local/bin/

# Copy builder files
COPY --from=builder /cryptopepe-reader/vendor/cryptopepe.io/cryptopepe-svg/builder/tmpl /app/tmpl
COPY --from=builder /cryptopepe-reader/vendor/cryptopepe.io/cryptopepe-svg/builder/builder.tmpl /app/builder.tmpl

# Copy bio builder file
COPY --from=builder /cryptopepe-reader/bio-gen/bio_config.yml /app/bio_config.yml

# Copy credentials file
COPY --from=builder /cryptopepe-reader/datastore-key.json /app/datastore-key.json

WORKDIR /app
ENTRYPOINT ["cryptopepe-reader"]
