

## Run service with docker-compose

Start `docker-compose up -d`

Rebuild `docker-compose build`

Stop `docker-compose down -v`

## use

GET plants  `curl localhost:8080/plant`

Create plant in postgres 

```
curl --location --request POST 'localhost:8080/plant' \
--header 'Content-Type: application/json' \
--data-raw '{
"ident":"111",
"name":"Car",
"status":"available",
"price":"4000"
}'

```

Create plant in mongodb 

```
curl --location --request POST 'localhost:8080/plantm' \
--header 'Content-Type: application/json' \
--data-raw '{
"ident":"111",
"name":"Car",
"status":"available",
"price":"4000"
}'

```


Check price for certain plant with ident

```
curl --location --request POST 'localhost:8080/price' --header 'Content-Type: application/json' --data-raw '{
"ident":"111"
}'


```





## Check docker-compose logs
All `docker-compose logs`

Individual service


`docker-compose logs bookmark-service`

`docker-compose logs postgres`
