#!/bin/bash

gql-query-from-file()
{
    local file=$1
    local gql=$(cat "$file")

    gql-query "$gql"
}

gql-query()
{
    local gql_string=$1

    query=$(jq -n \
        --arg query_gql "$gql_string" \
        '{ "query": $query_gql }' \
    )

    curl -k -XPOST -d "$query" https://[::1]:4343/graph | jq .
}

gql-query-from-file query.gql
gql-query-from-file mutate.gql

echo
