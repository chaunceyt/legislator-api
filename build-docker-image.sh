#!/bin/bash

APP_VERSION=$1
docker build -t chaunceyt/legislators-api -t chaunceyt/legislators-api:${APP_VERSION} .
docker push chaunceyt/legislators-api:${APP_VERSION}
docker push chaunceyt/legislators-api
