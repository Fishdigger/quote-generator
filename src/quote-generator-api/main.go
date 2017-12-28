package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(responseWriter http.ResponseWriter, request *http.Request) {
		responseWriter.Write([]byte("Hello, World!!! from api"))
	})

	http.Handle("/special", &helloWorldHandler{message: "Tenting this!!!"})

	http.ListenAndServe(":4000", nil)
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
