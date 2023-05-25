package env

import (
	"github.com/caarlos0/env/v8"
	"github.com/joho/godotenv"
	"log"
)

var loaded = false

type values struct {
	Env string `env:"ENV"`

	AppPort string `env:"APP_PORT" envDefault:"9000"`

	DbServer   string `env:"DB_SERVER" envDefault:"mongodb://127.0.0.1:27017"`
	DbName     string `env:"DB_NAME" envDefault:"book_app"`
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
		panic("Error loading .env file")
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
