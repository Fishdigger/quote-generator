package router

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/Fishdigger/quote-generator/src/services"
)

// Startup all routes in the api
func Startup() {
	r := mux.NewRouter()

	services.QuotesRegisterRoutes(r)

	http.Handle("/", r)

	// simple handle function example
	http.HandleFunc("/test2", func(responseWriter http.ResponseWriter, request *http.Request) {
		responseWriter.Write([]byte("Hello, World!!! from api"))
	})

	// template parsing handler example
	http.Handle("/special", &helloWorldHandler{message: "Tenting this!!!"})

	// simple handle function example
	http.HandleFunc("/test", testHandler)
}

func testHandler(responseWriter http.ResponseWriter, request *http.Request) {
	responseWriter.Write([]byte("Hello, World!!! from test function"))
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
