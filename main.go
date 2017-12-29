package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Fishdigger/quote-generator/src/router"
)

func main() {
	addr, err := getEnvPort()
	if err != nil {
		log.Fatal(err)
	}

	router.Startup()

	if err := http.ListenAndServe(addr, nil); err != nil {
		panic(err)
	}
}

func getEnvPort() (string, error) {
	port := os.Getenv("PORT")
	if port == "" {
		return "", fmt.Errorf("$PORT not set")
	}
	return ":" + port, nil
}
