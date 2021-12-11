package main

import (
	"fmt"
	"log"
	"net/http"
	"user-api/src/config"
	"user-api/src/router"
)

func main() {
	config.Load()
	fmt.Printf("API running on port: %d", config.Port)
	r := router.Generate()

	fmt.Print(config.SecretKey)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
