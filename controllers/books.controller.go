package books

import (
	"context"
	"errors"
	"fiber-book-app/helpers"
	"fiber-book-app/models"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

var Book = models.BooksCollection()

// AllBooks - Get all books
func AllBooks(c *fiber.Ctx) error {
	// get all books
	cursor, err := Book.Find(context.Background(), bson.M{}, models.ProjectAllBooks)
	if err != nil {
		return helpers.DbQueryError(c, err)
	}

	// get results
	var results = make([]bson.M, 0)
	if err = cursor.All(context.Background(), &results); err != nil {
		return helpers.DbQueryError(c, err)
	}

	return c.JSON(results)
}

// AddBook - Add new book
func AddBook(c *fiber.Ctx) error {
	// parse body
	body := new(models.BookForm)
	if err := c.BodyParser(body); err != nil {
		return helpers.BadRequest(c, err)
	}

	// validate body
	if ok, message := models.ValidateBookForm(*body); !ok {
		return helpers.BadRequest(c, errors.New(message))
	}

	now := primitive.NewDateTimeFromTime(time.Now())
	newBook := models.Book{
		ID:          primitive.NewObjectID(),
		Title:       body.Title,
		Description: body.Description,
		Available:   body.Available,
		UpdatedAt:   now,
		CreatedAt:   now,
	}

	// add book
	_, err := Book.InsertOne(context.Background(), newBook, nil)
	if err != nil {
		return helpers.DbQueryError(c, err)
	}

	return c.JSON(newBook)
}

// GetBook - Get single book
func GetBook(c *fiber.Ctx) error {
	// get bookId from locals
	bookId := c.Locals("bookId").(primitive.ObjectID)

	// find Book in db
	var book models.Book
	err := Book.FindOne(context.Background(), bson.M{"_id": bookId}).Decode(&book)

	if err != nil && err != mongo.ErrNoDocuments {
		return helpers.DbQueryError(c, err)
	}

	return c.JSON(book)
}

// UpdateBook - Update book
func UpdateBook(c *fiber.Ctx) error {
	// parse body
	body := new(models.BookForm)
	if err := c.BodyParser(body); err != nil {
		return helpers.BadRequest(c, err)
	}

	// validate body
	if ok, message := models.ValidateBookForm(*body); !ok {
		return helpers.BadRequest(c, errors.New(message))
	}

	// get bookId from locals
	bookId := c.Locals("bookId").(primitive.ObjectID)

	// update book
	_, err := Book.UpdateOne(context.Background(), bson.M{"_id": bookId}, bson.M{
		"$set": bson.M{
			"title":       body.Title,
			"description": body.Description,
			"available":   body.Available,
		},
	}, nil)

	if err != nil {
		return helpers.DbQueryError(c, err)
	}

	return c.JSON(fiber.Map{
		"message": "Book updated successfully!",
	})
}

// DeleteBook - Delete book
func DeleteBook(c *fiber.Ctx) error {
	// get bookId from locals
	bookId := c.Locals("bookId").(primitive.ObjectID)

	// delete book
	_, err := Book.DeleteOne(context.Background(), bson.M{"_id": bookId})

	if err != nil {
		return helpers.DbQueryError(c, err)
	}

	return c.JSON(fiber.Map{
		"message": "Book deleted successfully!",
	})
}

// DeleteAllBooks - Delete all books
func DeleteAllBooks(c *fiber.Ctx) error {
	deleted, err := Book.DeleteMany(context.Background(), bson.M{})

	if err != nil {
		return helpers.DbQueryError(c, err)
	}

	return c.JSON(fiber.Map{
		"message": fmt.Sprintf("Deleted %v books successfully!", deleted.DeletedCount),
	})
}
