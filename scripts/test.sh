#!/bin/bash

cd ../
grpcurl -plaintext -proto proto/bsbLookup.proto -d '{"bsb": "638010" }' localhost:8080 bsbLookup.bsbLookup.Validate