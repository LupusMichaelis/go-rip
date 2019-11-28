Playground to learn how to make an API, would be REST or GraphQL.

```
IMAGE_NAME=rip-api
docker build -t $RIPAPI .
CONTAINERID=$(docker run -d $RIPAPI)
CONTAINERIP=$(docker exec $CONTAINERID hostname -i)

docker logs -f $CONTAINERID
```
