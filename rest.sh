#!/bin/bash

curl -X GET -i -k https://[::1]:4343/rest/country/en

curl \
    -ik \
    -X POST \
    -H 'Content-Type: application/json; charset=UTF-8' \
    -d '{"code": "to", "Name": "Toto"}' \
    https://[::1]:4343/rest/country


curl \
    -ik \
    -X PUT \
    -H 'Content-Type: application/json; charset=UTF-8' \
    -d '{"code": "to", "Name": "Togo"}' \
    https://[::1]:4343/rest/country/to
