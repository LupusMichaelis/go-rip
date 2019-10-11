#!/bin/bash

gql=$(cat query.gql)
query=$(jq -n --arg gql "$gql" '{ "query": $gql }')

curl -k -XPOST -d "$query" https://[::1]:4343/graph | jq .
