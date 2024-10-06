package server

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
)

const envFile = ".env"

var EnvVars map[string]string

func Initialize() {
	err := godotenv.Load(envFile)
	if err != nil {
		panic("Error loading .env file")
	}

	EnvVars = make(map[string]string)

	for _, env := range os.Environ() {
		pair := strings.Split(env, "=")
		EnvVars[pair[0]] = pair[1]
	}
}