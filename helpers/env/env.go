package env

import (
	"fmt"
	"github.com/caarlos0/env/v8"
	"github.com/joho/godotenv"
	"log"
)

var loaded = false

type values struct {
	Env string `env:"ENV" envDefault:"development"`

	AppPort string `env:"APP_PORT" envDefault:"9000"`

	DbServer   string `env:"MONGODB_SERVER" envDefault:"mongodb://127.0.0.1:27017"`
	DbName     string `env:"MONGODB_DATABASE" envDefault:"books"`
	DbPassword string `env:"DB_PASSWORD" envDefault:""`
}

var Values values

func init() {
	LoadIfNotLoaded()
}

func LoadIfNotLoaded() {
	if loaded {
		return
	}

	// Load .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file, using default values")
	}

	Values = values{}
	if err := env.Parse(&Values); err != nil {
		log.Printf("%+v\n", err)
		panic("Error parsing env variables")
	}

	// mark as loaded
	loaded = true
}

func IsDevelopment() bool {
	return Values.Env == "development"
}
