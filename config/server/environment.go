package server

import (
	"github.com/joho/godotenv"
	"os"
)

var env = Environment{}

type Environment struct {
	JwtSecretKey string
	AppUrl       string
	Profile      string

	PgDbHost     string
	PgDbPort     string
	PgDbUser     string
	PgDbPassword string
	PgDbName     string

	RabbitmqHost     string
	RabbitmqPort     string
	RabbitmqUser     string
	RabbitmqPassword string
	RabbitmqVhost    string
}

func LoadEnvironment() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}

	loadEnvVars()
}

func GetEnvironment() Environment {
	return env
}

func loadEnvVars() {
	env.JwtSecretKey = panicIfEnvVarIsEmpty("JWT_SECRET_KEY")
	env.AppUrl = panicIfEnvVarIsEmpty("APP_URL")
	env.Profile = panicIfEnvVarIsEmpty("PROFILE")

	env.PgDbHost = panicIfEnvVarIsEmpty("PG_DB_HOST")
	env.PgDbPort = panicIfEnvVarIsEmpty("PG_DB_PORT")
	env.PgDbUser = panicIfEnvVarIsEmpty("PG_DB_USER")
	env.PgDbPassword = panicIfEnvVarIsEmpty("PG_DB_PASSWORD")
	env.PgDbName = panicIfEnvVarIsEmpty("PG_DB_NAME")

	env.RabbitmqHost = panicIfEnvVarIsEmpty("RABBITMQ_HOST")
	env.RabbitmqPort = panicIfEnvVarIsEmpty("RABBITMQ_PORT")
	env.RabbitmqUser = panicIfEnvVarIsEmpty("RABBITMQ_USER")
	env.RabbitmqPassword = panicIfEnvVarIsEmpty("RABBITMQ_PASSWORD")
	env.RabbitmqVhost = panicIfEnvVarIsEmpty("RABBITMQ_VHOST")
}

func panicIfEnvVarIsEmpty(key string) string {
	if key == "" {
		panic("Environment variable " + key + " is empty")
	}
	return os.Getenv(key)
}
