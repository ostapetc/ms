package env

import (
	log "github.com/sirupsen/logrus"
	"os"
)

func GetString(name string) *string {
	val := os.Getenv(name)

	if val == "" {
		log.Fatal("required env variable ", name)
	}

	return &val
}
