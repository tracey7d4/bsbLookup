#!/bin/bash

cd ../
grpcurl -plaintext -proto proto/bsbLookup.proto -d '{"bsb": "012030" }' localhost:8080 bsbLookup.bsbLookup.Validate