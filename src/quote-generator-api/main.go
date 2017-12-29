package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

func main() {
	addr, err := determineListenAddress()
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(responseWriter http.ResponseWriter, request *http.Request) {
		responseWriter.Write([]byte("Hello, World!!! from api"))
	})

	http.Handle("/special", &helloWorldHandler{message: "Tenting this!!!"})

	log.Printf("Listening on %s...\n", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		panic(err)
	}
}

type helloWorldHandler struct {
	message string
}

func (handler helloWorldHandler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	templatePattern := `Here is my content template being parsed <br /><br /><br />`

	t, err := template.New("first").Parse(templatePattern)
	if err != nil {
		fmt.Println("Problems parsing here!!!!", err)
	}
	err = t.Execute(responseWriter, nil)
	if err != nil {
		fmt.Println("Problems parsing here!!!!", err)
	}

	responseWriter.Write([]byte(`Hello World from struct!!!!` + handler.message))
}

func determineListenAddress() (string, error) {
	port := os.Getenv("PORT")
	if port == "" {
		return "", fmt.Errorf("$PORT not set")
	}
	return ":" + port, nil
}
