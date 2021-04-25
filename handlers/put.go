package handlers

import (
	"context"
	"log"
	"net/http"
	"encoding/json"

	"github.com/gorilla/mux"
	"github.com/imyashkale/mgo-crud/database"
	"github.com/imyashkale/mgo-crud/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdateBook(rw http.ResponseWriter, r *http.Request) {
		
	rw.Header().Set("Content-Type","application/json")

	var args = mux.Vars(r)

	id, err := primitive.ObjectIDFromHex(args["id"])
	if err != nil{
		log.Fatal(err)
	}

	var book models.Book

	//  Read all the params from body and conert to go's native type
	_ = json.NewDecoder(r.Body).Decode(&book)
	
	// filter query
	filter  := bson.M{"_id":id}

	// update query
	update := bson.D{
		{ "$set", bson.D{
				{"isbn" , book.Isbn},
				{"title", book.Title},
				{"author", bson.D{
					{"firstname", book.Author.FirstName},
					{"lastname", book.Author.LastName },
				}},
			}},
	}
	err = bookCollection.FindOneAndUpdate(context.TODO(),filter,update,).Decode(&book)

	if err != nil {
		database.GetError(err,rw)
		return
	};

	json.NewEncoder(rw).Encode(book);

}