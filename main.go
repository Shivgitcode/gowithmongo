package main

import (
	"fmt"
	"github/shivgitcode/gowithmongo/router"
	"log"
	"net/http"
)

func main() {

	r := router.Router()
	fmt.Println("Server is getting started...")
	log.Fatal(http.ListenAndServe(":4000", r))
	fmt.Println("Listening at Port 4000 ..")
}
