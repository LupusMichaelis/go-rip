#!/bin/bash

host='[::1]'
port=4343

host='172.17.0.2'
port=443

curl -X GET -i -k https://"$host":$port/r/country/en; echo
curl -X GET -i -k https://"$host":$port/r/country/uk; echo

curl \
    -ik \
    -X POST \
    -H 'Content-Type: application/json; charset=UTF-8' \
    -d '{"code": "to", "Name": "Toto"}' \
    https://"$host":$port/r/country
echo


curl \
    -ik \
    -X PUT \
    -H 'Content-Type: application/json; charset=UTF-8' \
    -d '{"code": "to", "Name": "Togo"}' \
    https://"$host":$port/r/country/to
echo
