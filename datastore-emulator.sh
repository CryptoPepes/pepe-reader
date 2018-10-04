#!/usr/bin/env bash

export CLOUDSDK_CORE_PROJECT=cryptopepe-192921
gcloud beta emulators datastore start --data-dir="$PWD/datastore-emu" --host-port="localhost:8081"
