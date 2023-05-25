package models

import (
	db "fiber-book-app/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Book struct {
	ID          primitive.ObjectID `bson:"_id" json:"id"`
	Title       string             `bson:"title" json:"title"`
	Description string             `bson:"description" json:"description"`
	Available   bool               `bson:"available" json:"available"`
	CreatedAt   primitive.DateTime `bson:"createdAt" json:"createdAt"`
	UpdatedAt   primitive.DateTime `bson:"updatedAt" json:"updatedAt"`
}

type BookForm struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Available   bool   `json:"available"`
}

var ProjectAllBooks = options.Find().SetProjection(bson.M{
	"_id":         0,
	"id":          "$_id",
	"title":       1,
	"description": 1,
	"available":   1,
	"updatedAt":   1,
	"createdAt":   1,
})

func BooksCollection() mongo.Collection {
	connectedDb := db.GetDb()
	collection := connectedDb.Collection("books")
	return *collection
}

func ValidateBookForm(book BookForm) (bool, string) {
	// validate book title
	if book.Title == "" {
		return false, "Title is required"
	}

	// validate book description
	if book.Description == "" {
		return false, "Description is required"
	}

	return true, ""
}
