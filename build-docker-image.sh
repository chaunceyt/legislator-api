#!/bin/bash

GOOS=linux GOARCH=amd64 go build -o build/legislators-api-linux cmd/legislator-query-api/*.go
cp data/legislators-current.json build/data/
cp data/legislators-district-offices.json build/data/
cp data/legislators-social-media.json build/data/
cp home.html build/
cp index.html build/

cd build
docker build -t chaunceyt/legislators-api:v0.0.3 .
docker push chaunceyt/legislators-api:v0.0.3
