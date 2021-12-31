package main

import (
	"log"
	"net/http"
	"project1/routes"
)

func main() {
	routes.InitApi()

	log.Fatal(http.ListenAndServe(":10000", routes.BaseRouter))
}
