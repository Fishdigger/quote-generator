package services

import (
	"net/http"

	"github.com/Fishdigger/quote-generator/src/database"
)

// QuoteRegisterRoutes register quotes for this service
func QuoteRegisterRoutes() {
	http.Handle("/get-quote", &getQuoteHandler{})
}

type getQuoteHandler struct {
	clientEmail string
}

func (handler getQuoteHandler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	dbSession := database.OpenSession()
	defer database.CloseSession(dbSession)

	responseWriter.Write([]byte(`Hello World from getQuote Service!!!!` + handler.clientEmail))
}
