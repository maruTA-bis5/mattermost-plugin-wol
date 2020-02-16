#!/bin/bash

docker build -t build-wol-plugin:latest .
mkdir -p .docker-cache/go
docker run --rm \
    -v "${PWD}:/repo" \
    -v "${HOME}/.docker-cache/go:/go" \
    -e "MM_SERVICESETTINGS_SITEURL=${MM_SERVICESETTINGS_SITEURL}" \
    -e "MM_ADMIN_USERNAME=${MM_ADMIN_USERNAME}" \
    -e "MM_ADMIN_PASSWORD=${MM_ADMIN_PASSWORD}" \
    -u "$(id -u $USER):$(id -g $USER)" build-wol-plugin:latest $@
