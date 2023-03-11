package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/MangeshSarjerao/MongoDbAPI/router"
)

func main() {
	fmt.Println("Welcome to build API using Mongo")
	r := router.Router()
	fmt.Println("Server is started....")
	log.Fatal(http.ListenAndServe(":8080", r))
	fmt.Println("Port is listening on 8080...")
}
