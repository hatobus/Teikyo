package util

import (
	"log"

	"github.com/joho/godotenv"
)

func Loadenv() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
