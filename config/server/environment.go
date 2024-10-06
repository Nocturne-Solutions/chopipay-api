package server

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
)

const envFile = ".env"

var EnvVars = make(map[string]string)

func LoadEnvirontment() {
	err := godotenv.Load(envFile)
	if err != nil {
		panic("Error loading .env file")
	}

	for _, env := range os.Environ() {
		pair := strings.Split(env, "=")
		EnvVars[pair[0]] = pair[1]
	}
}