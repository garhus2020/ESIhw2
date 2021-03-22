

# Documentation

In this document can be found some basic instructions needed to run and make use of this application. Is it important to mention that there is not any existing frontend consuming the API, so the application can be used only via HTTP requests (graphQL is implemented but not working).

### Deployment

Is it necessary first to have Docker installed in your local machine in case you desire to deploy the application in your computer. However, is it possible to access the application on our Hetzner virtual machine. [135.181.150.82:8080/plant ](http://135.181.150.82:8080/plant)

In case you wanted to deploy it locally, you should run the following commands:

```
docker-compose build
```

```
docker-compose up -d
```

Is it possible to stop the application running the following command:

```
docker-compose down -v
```

To get the docker-compose logs you should run:

```
docker-compose logs
```

To check the logs from postgres:

```
docker-compose logs postgres
```

## Application manual

This is a Go API which uses HTTP and graphQL. Unfortunately we didn't manage to make graphQL work but it's still possible to fully communicate with the application via HTML.

These are some basic CURL commands you can use to communicate with the server.

It's important to have in mind that the database table "plant" is split into two different databases (postgresr and mongodb). Different curl commands are used to get data from each database.

#### Get all plants from the database

```
curl localhost:8080/plant
```

#### Create one plant into the postgres repository

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

Create on plant into the mongodb repository

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

Check price for a certain plant with ID **(Notice ID is called "ident")**

```
curl --location --request POST 'localhost:8080/price' --header 'Content-Type: application/json' --data-raw '{
"ident":"111",
"start":"20",
"end":"22"
}'
```

Check status (availability) for certain plant with ID  **(Notice ID is called "ident")**

```
curl --location --request POST 'localhost:8080/status' --header 'Content-Type: application/json' --data-raw '{
"ident":"111",
"start":"20",
"end":"22"
}'
```

Retrieve all the cache 

```
curl localhost:8080/requests
```

