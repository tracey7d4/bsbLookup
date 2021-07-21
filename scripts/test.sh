#!/bin/bash

cd ../
grpcurl -plaintext -proto proto/bsbLookup.proto -d '{"bsb": "612312" }' localhost:8080 bsbLookup.bsbLookup.Validate