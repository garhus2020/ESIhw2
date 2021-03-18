package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	// SQL driver
	// https://www.calhoun.io/why-we-import-sql-drivers-with-the-blank-identifier/
	// The sql package must be used in conjunction with a database driver. In this case PostgreSQL driver.
	// See https://golang.org/s/sqldrivers for a list of drivers.
	_ "github.com/lib/pq"
)

const (
	httpServicePort    = 8080
	postgresConnection = "dbname=postgres host=postgres password=postgres user=postgres sslmode=disable port=5432"
	mongoConnection    = "mongodb://mongo:27017"
)

func main() {
	log.Println("Start plant server")

	// open Postgres connection
	dbConn, err := sql.Open("postgres", postgresConnection)
	if err != nil {
		log.Fatal(err)
	}

	//mongooooooooooooo
	dbConn2, err := mongo.NewClient(options.Client().ApplyURI(mongoConnection))
	if err != nil {
		log.Fatal(err)
	}
	err = dbConn2.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	plantRepository := NewPlantRepository(dbConn)
	plantmRepository := NewPlantmRepository(dbConn2)
	plantHandler := NewPlantHandler(plantmRepository, plantRepository)
	// mongooooooooooooooo

	// construct application

	router := mux.NewRouter()

	router.HandleFunc("/plant", plantHandler.CreatePlant).Methods(http.MethodPost)

	router.HandleFunc("/plantm", plantHandler.CreatePlantm).Methods(http.MethodPost)
	router.HandleFunc("/plant", plantHandler.GetPlants).Methods(http.MethodGet)

	router.HandleFunc("/price", plantHandler.GetPrice).Methods(http.MethodPost)

	// setup http server
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", httpServicePort),
		Handler: router,
	}

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatalf("Could not start server")
	}

	log.Println("Stop plant server")
}
