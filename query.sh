#!/bin/bash

gql-query-from-file()
{
    local file=$1
    local gql=$(cat "$file")

    local operationName=$2
    gql-query "$gql" "$operationName"
}

gql-query()
{
    local gql_string=$1
    local operationName=$2

    if [[ -z "$operationName" ]]
    then
        query=$(jq -n \
            --arg query_gql "$gql_string" \
            '{ "query": $query_gql }' \
        )
    else
        query=$(jq -n \
            --arg query_gql "$gql_string" \
            --arg query_operation "$operationName" \
            '{ "query": $query_gql, "operationName": $query_operation }' \
        )
    fi

    echo $query
}

gql-query-call()
{
    local query="$1"
    curl -k \
        -H 'Content-type: application/graphql; charset=UTF-8' \
        -d "$query" https://[::1]:4343/g
}

gql-query-call "$(gql-query-from-file query.gql)" #| jq .
echo
gql-query-call "$(gql-query-from-file mutate.gql)" #| jq .
echo

#query="$(gql-query-from-file query.gql read) $(gql-query-from-file mutate.gql write)"
#gql-query-call "$query" | jq .

echo
