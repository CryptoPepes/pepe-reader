# Build Geth in a stock Go builder container
FROM golang:1.9-alpine as builder

RUN apk add --no-cache make gcc musl-dev linux-headers

ADD . /cryptopepe-reader
RUN cd /cryptopepe-reader && build/env.sh go build -v -o build/cryptopepe-reader .

# Pull Geth into a second stage deploy alpine container
FROM alpine:latest

# Define that we want this docker build argument
ARG GOOGLE_DATASTORE_KEY
# Retrieve the build argument, insert it into the environment
ENV GOOGLE_DATASTORE_KEY=$GOOGLE_DATASTORE_KEY

RUN apk add --no-cache ca-certificates
COPY --from=builder /cryptopepe-reader/build/cryptopepe-reader /usr/local/bin/

# Copy bio builder file
COPY --from=builder /cryptopepe-reader/bio-gen/bio_config.yml /app/bio_config.yml

# Write credentials file (read from env, decode, tee into file)
RUN echo -n "$GOOGLE_DATASTORE_KEY" | base64 -d | tee /app/datastore-key.json

WORKDIR /app
ENTRYPOINT ["cryptopepe-reader"]
