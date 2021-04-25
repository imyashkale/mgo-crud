package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/imyashkale/mgo-crud/database"
	"github.com/imyashkale/mgo-crud/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var bookCollection = database.ConnectDB()

// GetBooks : Get All the books
func GetBooks(rw http.ResponseWriter, r *http.Request) {

	//setting content type
	rw.Header().Set("Content-Type", "application/json")
	//Book
	var books []models.Book

	cur, err := bookCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		database.GetError(err, rw)
		return
	}
	defer cur.Close(context.TODO())

	// Iterating over each record in cursor
	for cur.Next(context.TODO()) {
		// for decoding single record
		var book models.Book
		//mongodb response to the the go' native type
		err := cur.Decode(&book)
		if err != nil {
			log.Fatal(err)
		}
		//Adding book to the slice of books
		books = append(books, book)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
	json.NewEncoder(rw).Encode(books)
}

//GetBook : Get single book with id
func GetBook(rw http.ResponseWriter, request *http.Request) {
	// set Header Content-Type json
	rw.Header().Set("Content-Type", "application/json")

	// Book
	var book models.Book

	// All the passed params in url.
	var args = mux.Vars(request)
	log.Println("ARGS :", args)

	// Converting Id String to => Type Primitive.ObjectID
	id, err := primitive.ObjectIDFromHex(args["id"])
	if err != nil {
		log.Fatal(err)
		return
	}
	// filter book collection by Id
	filter := bson.M{"_id": id}
	// mongodb-qry result directly wriiten to the book struct
	err = bookCollection.FindOne(context.TODO(), filter).Decode(&book)
	if err != nil {
		database.GetError(err, rw)
		return
	}

	// Converting Mongodb-result to JSON
	json.NewEncoder(rw).Encode(book)
}

// CreateBook : Create New Book
func CreateBook(rw http.ResponseWriter, r *http.Request) {
	// set Header
	rw.Header().Set("Content-Type", "application/json")

	var book models.Book
	log.Println("CREATE BODY :", r.Body)
	//Decode body params
	json.NewDecoder(r.Body).Decode(&book)

	result, err := bookCollection.InsertOne(context.TODO(), book)

	if err != nil {
		database.GetError(err, rw)
		return
	}

	json.NewEncoder(rw).Encode(result)

}

// DeleteBook : Delete Book from  db with id
func DeleteBook(rw http.ResponseWriter, r *http.Request) {

	rw.Header().Set("Content-Type", "application/json")

	var args = mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(args["id"])

	filter := bson.M{"_id": id}

	deleteResult, err := bookCollection.DeleteOne(context.TODO(), filter)
	if err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(rw).Encode(deleteResult)
}
