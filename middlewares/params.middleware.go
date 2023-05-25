package middlewares

import (
	"context"
	"errors"
	"fiber-book-app/helpers"
	"fiber-book-app/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var bookCollection = models.BooksCollection()

func LoadBookParam(c *fiber.Ctx) error {
	// get "id" param
	id := c.Params("id")

	// convert string id to primitive.ObjectID
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return helpers.BadRequest(c, errors.New("invalid book ID"))
	}

	// check if book exists
	bookExists, err := bookCollection.CountDocuments(context.Background(), bson.M{"_id": oid})
	if err != nil {
		return helpers.DbQueryError(c, err)
	}

	if bookExists == 0 {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Book not found",
		})
	}

	// store object id in locals
	c.Locals("bookId", oid)

	return c.Next()
}
