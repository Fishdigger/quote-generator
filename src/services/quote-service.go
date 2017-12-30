package services

import (
	"encoding/json"
	"net/http"
	"strconv"

	"gopkg.in/mgo.v2"

	"github.com/Fishdigger/quote-generator/src/database"
	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

// COLLECTONNAME ... the name of the collection
const COLLECTONNAME = "Quotes"

// Quote ... struct to hold quotes data
type Quote struct {
	ID       bson.ObjectId `bson:"_id,omitempty"`
	Quote    string        `bson:"quote,omitempty"`
	ClientID int64         `bson:"clientID,omitempty"`
}

// QuotesRegisterRoutes ... Register routes for this service
func QuotesRegisterRoutes(router *mux.Router) {
	router.HandleFunc("/get-quote", getQuotes).Methods("GET")
	router.HandleFunc("/get-quote/{id}", getQuote)
	router.HandleFunc("/get-quote/{id}", getQuote).Methods("GET")
	router.HandleFunc("/get-quote/{id}", insertQuote).Methods("POST")
	router.HandleFunc("/get-quote/{id}", updateQuote).Methods("PUT")
	router.HandleFunc("/get-quote/{id}", deleteQuote).Methods("DELETE")
}

func getCollection() (dbSession *mgo.Session, collection *mgo.Collection) {
	dbSession = database.OpenSession()
	collection = dbSession.DB(database.DBNAME).C(COLLECTONNAME)
	return
}

func getQuotes(responseWriter http.ResponseWriter, request *http.Request) {
	dbSession, collection := getCollection()
	defer dbSession.Close()

	var results []Quote
	collection.Find(nil).All(&results)

	respondWithJSON(responseWriter, results)
}

func getQuote(responseWriter http.ResponseWriter, request *http.Request) {
	dbSession, collection := getCollection()
	defer dbSession.Close()

	params := mux.Vars(request)
	id, _ := strconv.ParseInt(params["id"], 10, 32)

	var result = Quote{}
	collection.Find(bson.M{"clientID": id}).One(&result)

	respondWithJSON(responseWriter, result)
}

func insertQuote(responseWriter http.ResponseWriter, request *http.Request) {
	dbSession, collection := getCollection()
	defer dbSession.Close()

	params := mux.Vars(request)
	var quote Quote
	_ = json.NewDecoder(request.Body).Decode(&quote)
	quote.ID = bson.ObjectIdHex(params["id"])

	_ = collection.Insert(quote)

	respondWithJSON(responseWriter, quote)
}

func updateQuote(responseWriter http.ResponseWriter, request *http.Request) {
	dbSession, collection := getCollection()
	defer dbSession.Close()

	params := mux.Vars(request)
	var quote Quote
	_ = json.NewDecoder(request.Body).Decode(&quote)
	quote.ID = bson.ObjectIdHex(params["id"])

	_ = collection.Update(quote.ID, quote)

	//change := bson.M{"$push": bson.M{"sections": bson.M{"name": "office"}}}
	//_ = collection.Update(quote.ID, change)

	respondWithJSON(responseWriter, quote)
}

func deleteQuote(responseWriter http.ResponseWriter, request *http.Request) {
	dbSession, collection := getCollection()
	defer dbSession.Close()

	params := mux.Vars(request)
	var quote Quote
	_ = json.NewDecoder(request.Body).Decode(&quote)
	quote.ID = bson.ObjectIdHex(params["id"])

	_ = collection.RemoveId(quote.ID)
	//_ = collection.Remove(quote.ID)

	respondWithJSON(responseWriter, quote)
}
