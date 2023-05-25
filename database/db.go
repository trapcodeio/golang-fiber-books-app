package db

import (
	"context"
	"fiber-book-app/helpers/env"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"strings"
)

// database - Hold the database connection
var database *mongo.Database

// Connected - Hold the connection status
var Connected bool = false

// ConnectToDatabase - Connect to the database
func ConnectToDatabase() {
	if Connected {
		return
	}

	// Load .env file
	env.LoadIfNotLoaded()

	// parse <dbname> and <password> from DB_SERVER
	DbServer := env.Values.DbServer
	if DbServer == "" {
		panic("DB_SERVER is not set")
	} else {
		// parse <dbname> and <password> from DB_SERVER
		DbServer = strings.Replace(DbServer, "<dbname>", env.Values.DbName, 1)
		DbServer = strings.Replace(DbServer, "<password>", env.Values.DbPassword, 1)
	}

	DbName := env.Values.DbName
	if DbName == "" {
		panic("DB_NAME is not set")
	}

	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(DbServer))
	if err != nil {
		panic(err)
	}

	database = client.Database(DbName)
	Connected = true

	fmt.Println("Connected to MongoDB: [" + DbName + "]")
}

// GetDb - Get db connection.
func GetDb() *mongo.Database {
	if !Connected {
		ConnectToDatabase()
	}

	return database
}
