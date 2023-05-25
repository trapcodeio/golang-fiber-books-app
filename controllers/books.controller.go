package books

import (
	"context"
	"errors"
	"fiber-book-app/helpers"
	"fiber-book-app/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

var Book = models.BooksCollection()

// AllBooks - Get all books
func AllBooks(c *fiber.Ctx) error {
	// get all books
	cursor, err := Book.Find(context.Background(), bson.M{})
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
	if ok, message := models.Validate(*body); !ok {
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
	return c.SendString("GetBook")
}

// UpdateBook - Update book
func UpdateBook(c *fiber.Ctx) error {
	return c.SendString("UpdateBook")
}
