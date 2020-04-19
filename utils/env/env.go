package env

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func GetString(name string, def string) *string {
	val := os.Getenv(name)

	if val != "" {
		return &val
	}

	if def == "" {
		log.Fatal("required env variable ", name)
	}

	return &def
}
