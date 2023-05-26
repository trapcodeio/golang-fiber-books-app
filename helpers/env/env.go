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

	DbServer   string `env:"DB_SERVER" envDefault:"mongodb://mongo/books"`
	DbName     string `env:"DB_NAME" envDefault:"books"`
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

	// print env values for debugging
	fmt.Println("Envs", Values)
}

func IsDevelopment() bool {
	return Values.Env == "development"
}
