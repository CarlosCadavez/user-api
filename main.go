package main

import (
	"fmt"
	"log"
	"net/http"
	"user-api/src/router"
)

func main() {
	fmt.Println("Running API")
	r := router.Generate()

	log.Fatal(http.ListenAndServe(":5000", r))
}
